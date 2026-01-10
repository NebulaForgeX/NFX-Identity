package messaging

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

// BusPublisher 消息总线发布器。
// 封装了消息发布器和 Exchange 解析器，提供消息键到 Exchange + RoutingKey 的转换。
type BusPublisher struct {
	pub              message.Publisher
	exchangeResolver *ExchangeResolver
}

// NewBusPublisher 创建新的消息总线发布器。
//
// 参数:
//   - pub: 底层消息发布器（如 RabbitMQ 发布器）
//   - exchangeResolver: Exchange 解析器，用于将消息键转换为 Exchange + RoutingKey
//
// 返回:
//   - *BusPublisher: 消息总线发布器实例
//
// 示例:
//
//	publisher := messaging.NewBusPublisher(rabbitmqPublisher, exchangeResolver)
func NewBusPublisher(
	pub message.Publisher,
	exchangeResolver *ExchangeResolver,
) *BusPublisher {
	return &BusPublisher{pub, exchangeResolver}
}

// Publish 发布消息到指定的 RoutingKey。
// 注意：对于 RabbitMQ，watermill-amqp 使用 topic 参数作为 RoutingKey
//
// 参数:
//   - routingKey: RoutingKey（实际的路由键，不是消息键）
//   - msgs: 要发布的消息列表
//
// 返回:
//   - error: 发布失败时返回错误
//
// 注意: 通常使用 PublishMessage 函数来发布类型化消息，而不是直接调用此方法
func (p *BusPublisher) Publish(routingKey string, msgs ...*message.Message) error {
	return p.pub.Publish(routingKey, msgs...)
}

// Close 关闭发布器。
// 释放所有资源并停止发布。
//
// 返回:
//   - error: 关闭失败时返回错误
func (p *BusPublisher) Close() error {
	return p.pub.Close()
}

// GetExchange 根据消息键获取 Exchange 和 RoutingKey。
//
// 参数:
//   - key: 消息键，例如 "user_created"
//
// 返回:
//   - ExchangeRouting: Exchange 和 RoutingKey，如果找到
//   - bool: 是否找到对应的 Exchange 和 RoutingKey
//
// 示例:
//
//	routing, ok := publisher.GetExchange("user_created")
//	if ok {
//	    fmt.Println("Exchange:", routing.Exchange, "RoutingKey:", routing.RoutingKey)
//	}
func (p *BusPublisher) GetExchange(key MessageKey) (ExchangeRouting, bool) {
	return p.exchangeResolver.GetExchange(key)
}
