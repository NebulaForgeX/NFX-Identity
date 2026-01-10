// Package messaging 提供类型安全的 RabbitMQ 消息传递功能。
//
// 核心功能:
//   - 类型化消息发布：基于泛型提供类型安全的消息发布
//   - 消息路由：根据消息类型自动路由到对应处理器（通过 Exchange/Queue）
//   - 路由管理：通过 ExchangeResolver 和 QueueResolver 管理消息路由
//   - 消息验证：支持 Validatable 接口进行消息数据验证
//
// # 创建消息的多种方式
//
// ## 方式一：自动生成 MessageType（推荐，最简洁）
//
// 只需要实现 RoutingKey() 方法，MessageType 会自动从类型名生成。
// 生成的规则：类型名去掉 "Message" 或 "Event" 后缀，转换为 snake_case。
//
//	package access
//
//	import "nfxid/pkgs/rabbitmqx/messaging"
//
//	// GrantsInvalidateCacheMessage 会自动生成 MessageType: "grants_invalidate_cache"
//	type GrantsInvalidateCacheMessage struct {
//	    ID string `json:"id"`
//	}
//
//	func (GrantsInvalidateCacheMessage) RoutingKey() messaging.MessageKey {
//	    return "access"
//	}
//
//	// 使用：直接创建结构体
//	msg := access.GrantsInvalidateCacheMessage{ID: "grant-123"}
//
// ## 方式二：自定义 MessageType
//
// 如果需要自定义 MessageType，实现 MessageType() 方法即可。
//
//	package custom
//
//	import "nfxid/pkgs/rabbitmqx/messaging"
//
//	type UserCreatedMessage struct {
//	    UserID string `json:"user_id"`
//	    Name   string `json:"name"`
//	}
//
//	// 自定义 MessageType
//	func (UserCreatedMessage) MessageType() messaging.MessageType {
//	    return "custom.user.created"  // 自定义的消息类型
//	}
//
//	func (UserCreatedMessage) RoutingKey() messaging.MessageKey {
//	    return "custom"
//	}
//
//	// 使用：直接创建结构体
//	msg := custom.UserCreatedMessage{
//	    UserID: "user-123",
//	    Name:   "John Doe",
//	}
//
// ## 方式三：使用 BaseMessage（可选，适用于需要动态设置 MessageType 的场景）
//
//	package example
//
//	import "nfxid/pkgs/rabbitmqx/messaging"
//
//	type DynamicMessage struct {
//	    messaging.BaseMessage
//	    Data map[string]interface{} `json:"data"`
//	}
//
//	// 使用构造函数创建
//	msg := DynamicMessage{
//	    BaseMessage: messaging.NewMessage(
//	        "example.dynamic.message",
//	        "example",
//	    ),
//	    Data: map[string]interface{}{"key": "value"},
//	}
//
// # 发布消息
//
//	// 创建发布器
//	publisher := messaging.NewBusPublisher(rabbitmqPublisher, exchangeResolver)
//
//	// 发布消息（自动生成 MessageType）
//	msg := access.GrantsInvalidateCacheMessage{ID: "grant-123"}
//	err := messaging.PublishMessage(ctx, publisher, msg)
//
//	// 发布消息（带元数据）
//	err := messaging.PublishMessage(ctx, publisher, msg,
//	    messaging.WithMetadata(map[string]string{
//	        "trace_id": "abc123",
//	        "user_id":  "user-456",
//	    }),
//	)
//
//	// 发布消息（自定义消息 ID）
//	err := messaging.PublishMessage(ctx, publisher, msg,
//	    messaging.WithMessageID("custom-message-id"),
//	)
//
// # 注册消息处理器
//
//	// 创建订阅器和路由器
//	subscriber := messaging.NewSubscriber(rabbitmqSubscriber, queueResolver)
//	router, _ := messaging.NewMessageRouter(subscriber, messaging.MessageRouterConfig{
//	    CloseTimeout: 30 * time.Second,
//	    Logger:       logger,
//	})
//
//	// 注册处理器（自动生成 MessageType 的消息）
//	messaging.RegisterHandler(router, func(ctx context.Context, msg access.GrantsInvalidateCacheMessage, rawMsg *message.Message) error {
//	    // 处理消息
//	    log.Printf("处理缓存清除消息: %s", msg.ID)
//	    return cache.Clear(msg.ID)
//	})
//
//	// 注册处理器（自定义 MessageType 的消息）
//	messaging.RegisterHandler(router, func(ctx context.Context, msg custom.UserCreatedMessage, rawMsg *message.Message) error {
//	    // 处理消息
//	    log.Printf("用户创建: %s - %s", msg.UserID, msg.Name)
//	    return nil
//	})
//
//	// 启动路由器
//	go func() {
//	    if err := router.Run(ctx); err != nil {
//	        log.Fatal(err)
//	    }
//	}()
//
// # MessageType 生成规则
//
// 当消息没有实现 MessageType() 方法时，会自动从类型名生成：
//
//   - 类型名: GrantsInvalidateCacheMessage
//   - 去掉 "Message" 或 "Event" 后缀: GrantsInvalidateCache
//   - 转换为 snake_case: grants_invalidate_cache
//
// 最终 MessageType: "grants_invalidate_cache"
package messaging

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

// MessageType 标识消息的唯一类型。
// 例如: "access.grants.invalidate_cache", "clients.apps.invalidate_cache"
type MessageType = string

// MessageHandler 处理特定类型的消息。
//
// 参数:
//   - ctx: 上下文，用于传递追踪信息、超时控制等
//   - msg: 消息对象，已反序列化的类型化消息
//   - rawMsg: 原始消息，包含元数据和载荷
//
// 返回:
//   - error: 处理错误，如果返回非 nil，消息将被标记为失败
type MessageHandler[T TypedMessage] func(ctx context.Context, msg T, rawMsg *message.Message) error

// Validatable 可验证接口。
// 实现此接口的消息可以在处理前进行验证。
type Validatable interface {
	// Validate 验证消息数据的有效性。
	// 如果消息数据无效，返回错误。
	Validate() error
}

// TypedMessage 类型化消息接口。
// 所有消息都必须实现此接口，提供路由键。
// 如果消息实现了 CustomTypedMessage 接口（包含 MessageType() 方法），则使用自定义的 MessageType。
// 否则，MessageType 会自动从类型名生成。
type TypedMessage interface {
	// RoutingKey 返回消息所属的路由键（MessageKey）。
	// 用于确定消息应该发布到哪个 Exchange/Queue。
	RoutingKey() MessageKey
}

// CustomTypedMessage 自定义类型化消息接口。
// 如果消息实现了此接口，MessageTypeOf 会使用自定义的 MessageType() 方法。
// 如果消息只实现了 TypedMessage（没有 MessageType()），则自动生成 MessageType。
type CustomTypedMessage interface {
	TypedMessage
	// MessageType 返回消息的类型标识符。
	// 用于在消息路由时识别消息类型。
	MessageType() MessageType
}

// MessageTypeOf 返回类型参数对应的 MessageType。
// 使用泛型在编译时确定消息类型，无需创建实例。
//
// 行为:
//  1. 首先检查消息是否实现了 CustomTypedMessage 接口（有 MessageType() 方法）
//  2. 如果实现了，使用自定义的 MessageType() 方法
//  3. 如果没有实现，自动使用 AutoMessageType 从类型名生成
//
// 示例:
//
//	// 方式一：用户实现 MessageType() 方法（自定义）
//	type MyMessage struct { ID string }
//	func (MyMessage) MessageType() MessageType { return "custom.message.type" }
//	func (MyMessage) RoutingKey() MessageKey { return "my_message" }
//	messageType := messaging.MessageTypeOf[MyMessage]()  // 返回: "custom.message.type"
//
//	// 方式二：不实现 MessageType()，自动生成（推荐，最简洁）
//	type MyMessage struct { ID string }
//	func (MyMessage) RoutingKey() MessageKey { return "my_message" }
//	// MessageType() 会自动使用 AutoMessageType 生成
//	messageType := messaging.MessageTypeOf[MyMessage]()  // 返回: "my_message"
func MessageTypeOf[T TypedMessage]() string {
	var zero T

	// 检查是否实现了 CustomTypedMessage 接口（有 MessageType() 方法）
	if customMessage, ok := any(zero).(CustomTypedMessage); ok {
		messageType := customMessage.MessageType()
		// 如果返回空字符串，也使用自动生成
		if messageType != "" {
			return messageType
		}
	}

	// 没有实现 MessageType() 或返回空字符串，自动生成
	return string(AutoMessageType[T]())
}

// MessageKeyOf 返回类型参数对应的 MessageKey。
// 使用泛型在编译时确定消息键，无需创建实例。
//
// 示例:
//
//	messageKey := messaging.MessageKeyOf[access.GrantsInvalidateCacheMessage]()
//	// 返回: "access"
func MessageKeyOf[T TypedMessage]() MessageKey {
	var zero T
	return zero.RoutingKey() // RoutingKey() 方法返回 MessageKey
}

// BaseMessage 消息的基础结构体。
// 提供一种可选的方式来实现 TypedMessage 接口。
// 消息可以直接实现 MessageType() 和 RoutingKey() 方法，也可以嵌入 BaseMessage。
type BaseMessage struct {
	messageType MessageType
	messageKey  MessageKey
}

// MessageType 返回消息类型。
// 实现 TypedMessage 接口的方法。
func (m BaseMessage) MessageType() MessageType {
	return m.messageType
}

// RoutingKey 返回消息键（MessageKey）。
// 实现 TypedMessage 接口的方法。
func (m BaseMessage) RoutingKey() MessageKey {
	return m.messageKey
}

// NewMessage 创建新消息。
// 提供一种可选的方式，类似于 Java 的构造函数，传入 messageType 和 messageKey 来创建基础消息。
//
// 参数:
//   - messageType: 消息类型，例如 "access.grants.invalidate_cache"
//   - messageKey: 消息键，例如 "access"
//
// 返回:
//   - BaseMessage: 实现了 TypedMessage 接口的基础消息
//
// 注意: 这是可选的方式之一。你也可以直接实现 MessageType() 和 RoutingKey() 方法，这样更简洁。
//
// 方式一 - 直接实现方法（推荐，更简洁）:
//
//	type GrantsInvalidateCacheMessage struct {
//	    ID string `json:"id"`
//	}
//
//	func (GrantsInvalidateCacheMessage) MessageType() MessageType {
//	    return "access.grants.invalidate_cache"
//	}
//	func (GrantsInvalidateCacheMessage) RoutingKey() MessageKey {
//	    return "access"
//	}
//
//	// 使用：直接创建结构体
//	msg := access.GrantsInvalidateCacheMessage{ID: "grant-123"}
//
// 方式二 - 使用 BaseMessage 和构造函数（可选）:
//
//	type GrantsInvalidateCacheMessage struct {
//	    messaging.BaseMessage
//	    ID string `json:"id"`
//	}
//
//	func NewGrantsInvalidateCacheMessage(id string) GrantsInvalidateCacheMessage {
//	    return GrantsInvalidateCacheMessage{
//	        BaseMessage: messaging.NewMessage(
//	            "access.grants.invalidate_cache",
//	            "access",
//	        ),
//	        ID: id,
//	    }
//	}
//
//	// 使用：通过构造函数创建
//	msg := access.NewGrantsInvalidateCacheMessage("grant-123")
func NewMessage(messageType MessageType, messageKey MessageKey) BaseMessage {
	return BaseMessage{
		messageType: messageType,
		messageKey:  messageKey,
	}
}
