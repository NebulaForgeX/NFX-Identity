package messaging

import (
	"fmt"
	"sync"

	"github.com/ThreeDotsLabs/watermill/message"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

// BusPublisher 消息总线发布器。
// 封装了消息发布器和 Exchange 解析器，提供消息键到 Exchange + RoutingKey 的转换。
// 支持在发送消息时动态指定 Exchange 类型。
type BusPublisher struct {
	pub              message.Publisher
	exchangeResolver *ExchangeResolver
	uri               string              // RabbitMQ URI，用于动态声明 Exchange
	exchangeConfig    ExchangeConfig      // Exchange 配置，用于动态声明
	declaredExchanges map[string]ExchangeType // 已声明的 Exchange（缓存）
	mu                sync.RWMutex         // 保护 declaredExchanges 的并发访问
}

// ExchangeConfig Exchange 配置，用于动态声明
type ExchangeConfig struct {
	Durable    bool
	AutoDelete bool
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
	return &BusPublisher{
		pub:              pub,
		exchangeResolver: exchangeResolver,
		declaredExchanges: make(map[string]ExchangeType),
	}
}

// NewBusPublisherWithConfig 创建新的消息总线发布器（带配置，支持动态声明 Exchange）。
//
// 参数:
//   - pub: 底层消息发布器（如 RabbitMQ 发布器）
//   - exchangeResolver: Exchange 解析器，用于将消息键转换为 Exchange + RoutingKey
//   - uri: RabbitMQ URI，用于动态声明 Exchange
//   - exchangeConfig: Exchange 配置（Durable, AutoDelete）
//
// 返回:
//   - *BusPublisher: 消息总线发布器实例
//
// 示例:
//
//	publisher := messaging.NewBusPublisherWithConfig(
//	    rabbitmqPublisher,
//	    exchangeResolver,
//	    "amqp://guest:guest@localhost:5672/",
//	    messaging.ExchangeConfig{Durable: true, AutoDelete: false},
//	)
func NewBusPublisherWithConfig(
	pub message.Publisher,
	exchangeResolver *ExchangeResolver,
	uri string,
	exchangeConfig ExchangeConfig,
) *BusPublisher {
	return &BusPublisher{
		pub:              pub,
		exchangeResolver: exchangeResolver,
		uri:              uri,
		exchangeConfig:   exchangeConfig,
		declaredExchanges: make(map[string]ExchangeType),
	}
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

// ensureExchange 确保 Exchange 存在且类型正确（如果指定了类型）。
// 如果 Exchange 不存在，自动创建；如果类型不匹配，返回错误。
func (p *BusPublisher) ensureExchange(exchangeName string, exchangeType ExchangeType) error {
	if exchangeType == "" {
		// 未指定类型，不需要检查
		return nil
	}

	// 检查是否已声明过
	p.mu.RLock()
	if declaredType, ok := p.declaredExchanges[exchangeName]; ok {
		p.mu.RUnlock()
		if declaredType != exchangeType {
			return fmt.Errorf("exchange %s already declared with type %s, cannot use type %s", exchangeName, declaredType, exchangeType)
		}
		return nil
	}
	p.mu.RUnlock()

	// 如果 URI 未设置，无法动态声明
	if p.uri == "" {
		// 没有 URI，无法动态声明，依赖预先声明
		return nil
	}

	// 声明 Exchange
	if err := p.declareExchange(exchangeName, exchangeType); err != nil {
		return err
	}

	// 缓存已声明的 Exchange
	p.mu.Lock()
	p.declaredExchanges[exchangeName] = exchangeType
	p.mu.Unlock()

	return nil
}

// SetDeclaredExchange 设置已声明的 Exchange（用于缓存）。
// 在预先声明 Exchange 后调用此方法，避免重复声明。
func (p *BusPublisher) SetDeclaredExchange(name string, exchangeType ExchangeType) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.declaredExchanges[name] = exchangeType
}

// declareExchange 声明 Exchange
func (p *BusPublisher) declareExchange(name string, exchangeType ExchangeType) error {
	conn, err := amqp091.Dial(p.uri)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ for exchange declaration: %w", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel for exchange declaration: %w", err)
	}
	defer ch.Close()

	arguments := make(amqp091.Table)
	if exchangeType == ExchangeTypeDelayedMessage {
		arguments["x-delayed-type"] = DefaultExchangeType.String()
	}

	err = ch.ExchangeDeclare(
		name,
		exchangeType.String(),
		p.exchangeConfig.Durable,
		p.exchangeConfig.AutoDelete,
		false, // internal
		false, // no-wait
		arguments,
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange %s (type: %s): %w", name, exchangeType, err)
	}

	return nil
}
