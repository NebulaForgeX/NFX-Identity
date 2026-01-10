package rabbitmqx

import (
	"fmt"
	"nfxid/pkgs/logx"
	"nfxid/pkgs/rabbitmqx/messaging"

	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

func NewPublisher(cfg *Config) (*messaging.BusPublisher, error) {
	amqpConfig := BuildAMQPConfig(cfg)

	defaultExchange := cfg.Exchange.Name
	if defaultExchange == "" {
		defaultExchange = "default"
	}
	defaultExchangeType := cfg.Exchange.Type
	if defaultExchangeType == "" {
		defaultExchangeType = messaging.DefaultExchangeType
	}
	logx.S().Infof("🔄 Initializing RabbitMQ Publisher with URI: %s (default exchange: %s, type: %s)", maskURI(cfg.URI), defaultExchange, defaultExchangeType.String())

	// 使用支持优先级的 Marshaler（总是启用，以便支持 WithPriority 选项）
	amqpConfig.Marshaler = &PriorityMarshaler{}

	pub, err := amqp.NewPublisher(amqpConfig, logx.NewZapWatermillLogger(logx.L()))
	if err != nil {
		logx.S().Errorf("❌ Failed to create RabbitMQ Publisher: %v", err)
		return nil, err
	}

	// ✅ 为每个不同的 Exchange 名称和类型组合声明 Exchange
	// 收集所有需要的 Exchange
	exchanges := make(map[string]messaging.ExchangeType) // exchange name -> exchange type
	for _, routing := range cfg.ProducerExchanges {
		exchange := routing.Exchange
		if exchange == "" {
			exchange = defaultExchange
		}
		exchangeType := routing.Type
		if exchangeType == "" {
			exchangeType = defaultExchangeType
		}
		exchanges[exchange] = exchangeType
	}

	// 声明所有需要的 Exchange（在创建 Publisher 之前）
	if err := declareExchanges(nil, exchanges, cfg.Exchange.Durable, cfg.Exchange.AutoDelete, cfg.URI); err != nil {
		logx.S().Errorf("❌ Failed to declare exchanges: %v", err)
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

	// ✅ 创建 BusPublisher，支持动态声明 Exchange
	busPublisher := messaging.NewBusPublisherWithConfig(
		pub,
		exchangeResolver,
		cfg.URI,
		messaging.ExchangeConfig{
			Durable:    cfg.Exchange.Durable,
			AutoDelete: cfg.Exchange.AutoDelete,
		},
	)

	// 预先声明配置中的 Exchange（可选，用于提前验证）
	if len(exchanges) > 0 {
		if err := declareExchanges(nil, exchanges, cfg.Exchange.Durable, cfg.Exchange.AutoDelete, cfg.URI); err != nil {
			logx.S().Warnf("⚠️ Failed to pre-declare exchanges (will be declared on-demand): %v", err)
			// 不返回错误，允许在发送消息时动态声明
		} else {
			// 将已声明的 Exchange 添加到缓存
			for name, exchangeType := range exchanges {
				busPublisher.SetDeclaredExchange(name, exchangeType)
			}
		}
	}

	logx.S().Infof("✅ Successfully connected to RabbitMQ Publisher: %s (exchanges: %d)", maskURI(cfg.URI), len(exchanges))
	return busPublisher, nil
}

// declareExchanges 声明所有需要的 Exchange
// 使用 amqp091-go 直接连接并声明，因为 watermill-amqp 的 Publisher 使用全局 Exchange 类型
// 这样可以为每个 Exchange 指定不同的类型（topic, fanout, direct, headers, x-delayed-message 等）
func declareExchanges(_ *amqp.Publisher, exchanges map[string]messaging.ExchangeType, durable, autoDelete bool, uri string) error {
	if len(exchanges) == 0 {
		return nil
	}

	// 创建临时连接来声明 Exchange
	conn, err := amqp091.Dial(uri)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ for exchange declaration: %w", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel for exchange declaration: %w", err)
	}
	defer ch.Close()

	// 声明所有需要的 Exchange
	for name, exchangeType := range exchanges {
		// 对于插件类型（如 x-delayed-message），可能需要额外的 arguments
		arguments := make(amqp091.Table)
		
		// x-delayed-message 需要指定 x-delayed-type
		if exchangeType == messaging.ExchangeTypeDelayedMessage {
			// 默认使用 topic 作为底层类型，可以通过配置覆盖
			arguments["x-delayed-type"] = messaging.DefaultExchangeType.String()
		}
		
		err := ch.ExchangeDeclare(
			name,                    // name
			exchangeType.String(),   // type
			durable,                 // durable
			autoDelete,              // auto-deleted
			false,                   // internal
			false,                   // no-wait
			arguments,               // arguments（用于插件类型）
		)
		if err != nil {
			return fmt.Errorf("failed to declare exchange %s (type: %s): %w", name, exchangeType, err)
		}
		logx.S().Infof("✅ Declared exchange: %s (type: %s, durable: %v)", name, exchangeType, durable)
	}

	return nil
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
