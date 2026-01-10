package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

// MessageRouterConfig 消息路由器配置。
type MessageRouterConfig struct {
	CloseTimeout           time.Duration           // 关闭超时时间，默认 10 秒
	Logger                 watermill.LoggerAdapter // 日志适配器，默认 NopLogger
	MessageTypeMetadataKey string                  // 消息类型元数据键，默认 "message_type"
}

func (c *MessageRouterConfig) setDefaults() {
	if c.Logger == nil {
		c.Logger = watermill.NopLogger{}
	}
	if c.CloseTimeout == 0 {
		c.CloseTimeout = 10 * time.Second
	}
	if c.MessageTypeMetadataKey == "" {
		c.MessageTypeMetadataKey = DefaultMessageTypeHeaderKey
	}
}

// MessageRouter 消息路由器。
// 根据消息类型自动路由消息到对应的处理器。
type MessageRouter struct {
	*message.Router
	sub             *BusSubscriber
	cfg             *MessageRouterConfig
	handlersByTopic map[string]map[string]message.NoPublishHandlerFunc
	muxInstalled    map[string]struct{}
}

// NewMessageRouter 创建新的消息路由器。
//
// 参数:
//   - sub: 消息总线订阅器
//   - cfg: 路由器配置
//
// 返回:
//   - *MessageRouter: 消息路由器实例
//   - error: 创建失败时返回错误
//
// 示例:
//
//	router, err := messaging.NewMessageRouter(subscriber, messaging.MessageRouterConfig{
//	    CloseTimeout: 30 * time.Second,
//	    Logger:       logger,
//	})
func NewMessageRouter(sub *BusSubscriber, cfg MessageRouterConfig) (*MessageRouter, error) {
	cfg.setDefaults()
	r, err := message.NewRouter(message.RouterConfig{CloseTimeout: cfg.CloseTimeout}, cfg.Logger)
	if err != nil {
		return nil, err
	}
	return &MessageRouter{
		Router:          r,
		sub:             sub,
		cfg:             &cfg,
		handlersByTopic: make(map[string]map[string]message.NoPublishHandlerFunc),
		muxInstalled:    make(map[string]struct{}),
	}, nil
}

// RegisterHandler 注册消息处理器。
// 根据消息的类型参数自动确定队列和消息类型，并注册对应的处理器。
//
// 参数:
//   - mr: 消息路由器
//   - handler: 消息处理器函数
//
// 注意: 如果消息键未找到或处理器已存在，会触发 panic
//
// 示例:
//
//	// 注册 GrantsInvalidateCacheMessage 的处理器
//	messaging.RegisterHandler(router, func(ctx context.Context, msg access.GrantsInvalidateCacheMessage, rawMsg *message.Message) error {
//	    log.Info("Received grant cache invalidation message", "id", msg.ID)
//	    // 处理消息逻辑
//	    return nil
//	})
func RegisterHandler[T TypedMessage](
	mr *MessageRouter,
	handler MessageHandler[T],
) {
	// Make sure the queue is bound to the message
	queueBinding, ok := mr.sub.GetQueue(MessageKeyOf[T]())
	if !ok || queueBinding.Queue == "" {
		var zero T
		panic(fmt.Sprintf("messaging: queue not found for message type %T", zero))
	}

	messageType := MessageTypeOf[T]()

	if mr.handlersByTopic[queueBinding.Queue] == nil {
		mr.handlersByTopic[queueBinding.Queue] = make(map[string]message.NoPublishHandlerFunc)
	}
	if _, exists := mr.handlersByTopic[queueBinding.Queue][messageType]; exists {
		panic(fmt.Sprintf("messaging: duplicate handler for %s on queue %s", messageType, queueBinding.Queue))
	}

	mr.handlersByTopic[queueBinding.Queue][messageType] = wrapHandler(handler)
	mr.ensureMux(queueBinding.Queue)
}

func (mr *MessageRouter) ensureMux(queue string) {
	if _, ok := mr.muxInstalled[queue]; ok {
		return
	}
	handlerName := fmt.Sprintf("%s__domain_mux", queue)
	mr.Router.AddConsumerHandler(
		handlerName,
		queue, // 对于 RabbitMQ，watermill-amqp 使用 topic 参数作为 Queue 名称
		mr.sub,
		func(msg *message.Message) error {
			mt := msg.Metadata.Get(mr.cfg.MessageTypeMetadataKey)
			if mt == "" {
				mr.Router.Logger().Error(fmt.Sprintf("missing %s header", mr.cfg.MessageTypeMetadataKey), nil, watermill.LogFields{
					"queue": queue, "uuid": msg.UUID,
				})
				return nil
			}
			m := mr.handlersByTopic[queue]
			if m == nil {
				mr.Router.Logger().Error(fmt.Sprintf("no handler map for queue %s", queue), nil, watermill.LogFields{
					"queue": queue, mr.cfg.MessageTypeMetadataKey: mt, "uuid": msg.UUID,
				})
				return nil
			}
			h, ok := m[mt]
			if !ok {
				mr.Router.Logger().Info(fmt.Sprintf("skip unknown %s", mr.cfg.MessageTypeMetadataKey), watermill.LogFields{
					"queue": queue, mr.cfg.MessageTypeMetadataKey: mt, "uuid": msg.UUID,
				})
				return nil
			}
			return h(msg)
		},
	)
	mr.muxInstalled[queue] = struct{}{}
}

func wrapHandler[T TypedMessage](msgHandler MessageHandler[T]) message.NoPublishHandlerFunc {
	return func(rawMsg *message.Message) error {
		var msg T
		if err := json.Unmarshal(rawMsg.Payload, &msg); err != nil {
			return fmt.Errorf("messaging: json unmarshal failed: %w", err)
		}

		if v, ok := any(&msg).(Validatable); ok {
			if err := v.Validate(); err != nil {
				return fmt.Errorf("messaging: validation failed: %w", err)
			}
		}

		// Pass context, for tracing / timeout control
		ctx := rawMsg.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		return msgHandler(ctx, msg, rawMsg)
	}
}
