package messaging

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

// BusSubscriber 消息总线订阅器。
// 封装了消息订阅器和 Queue 解析器，提供消息键到 Queue + BindingKey 的转换。
type BusSubscriber struct {
	sub           message.Subscriber
	queueResolver *QueueResolver
}

// NewSubscriber 创建新的消息总线订阅器。
//
// 参数:
//   - sub: 底层消息订阅器（如 RabbitMQ 订阅器）
//   - queueResolver: Queue 解析器，用于将消息键转换为 Queue + BindingKey
//
// 返回:
//   - *BusSubscriber: 消息总线订阅器实例
//
// 示例:
//
//	subscriber := messaging.NewSubscriber(rabbitmqSubscriber, queueResolver)
func NewSubscriber(
	sub message.Subscriber,
	queueResolver *QueueResolver,
) *BusSubscriber {
	return &BusSubscriber{sub, queueResolver}
}

// Subscribe 订阅指定 Queue 的消息。
// 注意：对于 RabbitMQ，watermill-amqp 使用 topic 参数作为 Queue 名称
//
// 参数:
//   - ctx: 上下文，用于控制订阅的生命周期
//   - queue: Queue 名称（实际的队列名称，不是消息键）
//
// 返回:
//   - <-chan *message.Message: 消息通道
//   - error: 订阅失败时返回错误
//
// 示例:
//
//	messages, err := subscriber.Subscribe(ctx, "user-queue")
//	if err != nil {
//	    return err
//	}
//	for msg := range messages {
//	    // 处理消息
//	}
func (s *BusSubscriber) Subscribe(ctx context.Context, queue string) (<-chan *message.Message, error) {
	return s.sub.Subscribe(ctx, queue)
}

// Close 关闭订阅器。
// 释放所有资源并停止订阅。
//
// 返回:
//   - error: 关闭失败时返回错误
func (s *BusSubscriber) Close() error {
	return s.sub.Close()
}

// GetQueue 根据消息键获取 Queue 和 BindingKey。
//
// 参数:
//   - key: 消息键，例如 "user_created"
//
// 返回:
//   - QueueBinding: Queue 和 BindingKey，如果找到
//   - bool: 是否找到对应的 Queue 和 BindingKey
//
// 示例:
//
//	binding, ok := subscriber.GetQueue("user_created")
//	if ok {
//	    fmt.Println("Queue:", binding.Queue, "BindingKey:", binding.BindingKey)
//	}
func (s *BusSubscriber) GetQueue(key MessageKey) (QueueBinding, bool) {
	return s.queueResolver.GetQueue(key)
}
