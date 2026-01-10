package messaging

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

// DefaultMessageTypeHeaderKey 默认的消息类型元数据键。
// 用于在消息元数据中存储消息类型。
const DefaultMessageTypeHeaderKey = "message_type"

// PublishOption 发布选项函数类型。
// 用于配置消息发布时的各种选项。
type PublishOption func(*message.Message)

// WithMetadata 设置消息的元数据。
// 将提供的键值对添加到消息的元数据中。
//
// 参数:
//   - meta: 元数据键值对映射
//
// 示例:
//
//	messaging.PublishMessage(ctx, publisher, msg,
//	    messaging.WithMetadata(map[string]string{
//	        "trace_id": "abc123",
//	        "user_id":  "user-456",
//	    }),
//	)
func WithMetadata(meta map[string]string) PublishOption {
	return func(m *message.Message) {
		for k, v := range meta {
			m.Metadata.Set(k, v)
		}
	}
}

// WithMessageID 设置消息的 ID。
// 如果不设置，将自动生成 UUID v7。
//
// 参数:
//   - id: 消息 ID 字符串
//
// 示例:
//
//	messaging.PublishMessage(ctx, publisher, msg,
//	    messaging.WithMessageID("custom-message-id-123"),
//	)
func WithMessageID(id string) PublishOption {
	return func(m *message.Message) { m.UUID = id }
}

// WithPartitionKey 设置消息的分区键。
// 用于控制消息路由到哪个分区。
//
// 参数:
//   - pk: 分区键字符串
//
// 示例:
//
//	messaging.PublishMessage(ctx, publisher, msg,
//	    messaging.WithPartitionKey("tenant-123"),
//	)
func WithPartitionKey(pk string) PublishOption {
	return func(m *message.Message) { m.Metadata.Set("partition_key", pk) }
}

// WithMessageTypeHeaderKey 设置消息类型头键的自定义名称。
// 默认使用 DefaultMessageTypeHeaderKey ("message_type")。
//
// 参数:
//   - key: 自定义的消息类型头键名称
//
// 示例:
//
//	messaging.PublishMessage(ctx, publisher, msg,
//	    messaging.WithMessageTypeHeaderKey("x-message-type"),
//	)
func WithMessageTypeHeaderKey(key string) PublishOption {
	return func(m *message.Message) { m.Metadata.Set("__message_type_header_key__", key) }
}

// WithExchangeType 设置 Exchange 类型（在发送消息时动态指定）。
// 如果 Exchange 不存在，会自动创建（使用指定的类型）。
// 如果 Exchange 已存在但类型不匹配，会返回错误。
//
// 参数:
//   - exchangeType: Exchange 类型（direct, topic, fanout, headers, x-delayed-message 等）
//
// 示例:
//
//	// 发送到 Fanout Exchange（广播）
//	messaging.PublishMessage(ctx, publisher, msg,
//	    messaging.WithExchangeType(messaging.ExchangeTypeFanout),
//	)
//
//	// 发送到延迟消息 Exchange（需要插件）
//	messaging.PublishMessage(ctx, publisher, msg,
//	    messaging.WithExchangeType(messaging.ExchangeTypeDelayedMessage),
//	    messaging.WithMetadata(map[string]string{"x-delay": "5000"}),
//	)
func WithExchangeType(exchangeType ExchangeType) PublishOption {
	return func(m *message.Message) {
		m.Metadata.Set("__exchange_type__", exchangeType.String())
	}
}

// PublishMessage 发布类型化消息到消息总线。
// 自动序列化消息、设置消息类型头、生成消息 ID，并发布到对应的 Exchange。
//
// 参数:
//   - ctx: 上下文，用于传递追踪信息等
//   - bp: 消息总线发布器
//   - msg: 要发布的消息对象，必须实现 TypedMessage 接口
//   - opts: 可选的发布选项，如 WithMetadata、WithMessageID 等
//
// 返回:
//   - error: 发布失败时返回错误
//
// 示例:
//
//	// 创建消息
//	msg := access.NewGrantsInvalidateCacheMessage("grant-123")
//
//	// 发布消息
//	err := messaging.PublishMessage(ctx, publisher, msg)
//	if err != nil {
//	    log.Error("failed to publish message", err)
//	}
//
//	// 带选项发布
//	err := messaging.PublishMessage(ctx, publisher, msg,
//	    messaging.WithMetadata(map[string]string{"trace_id": "abc123"}),
//	    messaging.WithPartitionKey("tenant-123"),
//	)
func PublishMessage[T TypedMessage](
	ctx context.Context,
	bp *BusPublisher,
	msg T,
	opts ...PublishOption,
) error {
	routing, ok := bp.GetExchange(MessageKeyOf[T]())
	if !ok || routing.RoutingKey == "" {
		return fmt.Errorf("messaging: exchange/routing_key not found for message type %T", msg)
	}

	payload, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("messaging: marshal failed: %w", err)
	}

	id, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("messaging: generate id failed: %w", err)
	}

	rawMsg := message.NewMessageWithContext(ctx, id.String(), payload)
	for _, opt := range opts {
		opt(rawMsg)
	}

	// ✅ 检查是否在发送消息时指定了 Exchange 类型
	if exchangeTypeStr := rawMsg.Metadata.Get("__exchange_type__"); exchangeTypeStr != "" {
		exchangeType := ExchangeType(exchangeTypeStr)
		if !exchangeType.IsValid() {
			return fmt.Errorf("messaging: invalid exchange type: %s", exchangeType)
		}
		// 确保 Exchange 存在且类型正确
		if err := bp.ensureExchange(routing.Exchange, exchangeType); err != nil {
			return fmt.Errorf("messaging: failed to ensure exchange %s (type: %s): %w", routing.Exchange, exchangeType, err)
		}
	}

	headerKey := rawMsg.Metadata.Get("__message_type_header_key__")
	if headerKey == "" {
		headerKey = DefaultMessageTypeHeaderKey
	}

	if rawMsg.Metadata.Get(headerKey) == "" {
		rawMsg.Metadata.Set(headerKey, MessageTypeOf[T]())
	}

	// 对于 RabbitMQ，watermill-amqp 使用 topic 参数作为 RoutingKey
	// Exchange 在配置中已经设置
	// 注意：消息优先级通过 WithPriority 选项设置，或通过消息元数据中的 "x-priority" 键
	return bp.Publish(routing.RoutingKey, rawMsg)
}
