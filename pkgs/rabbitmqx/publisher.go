package rabbitmqx

import (
	"nfxid/pkgs/logx"
	"nfxid/pkgs/rabbitmqx/messaging"

	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
)

func NewPublisher(cfg *Config) (*messaging.BusPublisher, error) {
	amqpConfig := BuildAMQPConfig(cfg)

	defaultExchange := cfg.Exchange.Name
	if defaultExchange == "" {
		defaultExchange = "default"
	}
	logx.S().Infof("🔄 Initializing RabbitMQ Publisher with URI: %s (default exchange: %s)", maskURI(cfg.URI), defaultExchange)

	// 使用支持优先级的 Marshaler（总是启用，以便支持 WithPriority 选项）
	amqpConfig.Marshaler = &PriorityMarshaler{}

	pub, err := amqp.NewPublisher(amqpConfig, logx.NewZapWatermillLogger(logx.L()))
	if err != nil {
		logx.S().Errorf("❌ Failed to create RabbitMQ Publisher: %v", err)
		return nil, err
	}

	// 将 ProducerExchanges 转换为 ExchangeResolver 需要的格式
	keyToRouting := make(map[messaging.MessageKey]messaging.ExchangeRouting)
	for eventKey, routing := range cfg.ProducerExchanges {
		exchange := routing.Exchange
		if exchange == "" {
			exchange = defaultExchange
		}
		routingKey := routing.RoutingKey
		if routingKey == "" {
			routingKey = eventKey // 默认使用事件键作为 RoutingKey
		}
		keyToRouting[messaging.MessageKey(eventKey)] = messaging.ExchangeRouting{
			Exchange:   exchange,
			RoutingKey: routingKey,
		}
	}

	exchangeResolver, err := messaging.NewExchangeResolver(keyToRouting)
	if err != nil {
		logx.S().Errorf("❌ Failed to create exchange resolver: %v", err)
		return nil, err
	}

	logx.S().Infof("✅ Successfully connected to RabbitMQ Publisher: %s (exchange: %s)", maskURI(cfg.URI), defaultExchange)
	return messaging.NewBusPublisher(pub, exchangeResolver), nil
}

// maskURI 隐藏 URI 中的密码，用于日志输出
func maskURI(uri string) string {
	// 简单实现：将 amqp://user:password@host:port 转换为 amqp://user:***@host:port
	// 这里可以做得更复杂，但为了简单起见，只处理基本格式
	if len(uri) < 10 {
		return "***"
	}
	// 查找 @ 符号的位置
	for i := 8; i < len(uri); i++ {
		if uri[i] == '@' {
			// 找到密码部分，用 *** 替换
			for j := i - 1; j > 0; j-- {
				if uri[j] == ':' {
					return uri[:j+1] + "***" + uri[i:]
				}
			}
			break
		}
	}
	return uri
}
