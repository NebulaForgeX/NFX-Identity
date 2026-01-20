package rabbitmqx

import (
	"crypto/tls"
	"time"

	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	amqp091 "github.com/rabbitmq/amqp091-go"
	"nfxid/pkgs/rabbitmqx/messaging"
)

// BuildAMQPConfig 构建 AMQP 配置。
// 根据提供的 Config 创建适合 RabbitMQ 的 AMQP 配置。
func BuildAMQPConfig(c *Config) (amqp.Config, error) {
	uri, err := c.BuildURI()
	if err != nil {
		return amqp.Config{}, err
	}

	cfg := amqp.NewDurableQueueConfig(uri)

	// 配置连接选项（TLS 和 Reconnect）
	ConfigureConnection(&cfg, c.Connection)

	// 配置发布者选项
	ConfigurePublisher(&cfg, c.Producer)

	// 配置消费者选项
	ConfigureConsumer(&cfg, c.Consumer)

	// 配置交换机
	ConfigureExchange(&cfg, c.Exchange)

	// 配置队列（RabbitMQ 特有）
	ConfigureQueue(&cfg, c.Queue, c.Consumer)

	// 配置队列绑定（RabbitMQ 特有）
	ConfigureQueueBind(&cfg, c.QueueBind)

	return cfg, nil
}

// ConfigureConnection 配置连接相关选项（TLS、Reconnect 和底层 AMQP）。
// 如果配置项未启用或为空，则使用 watermill-amqp 的默认值。
func ConfigureConnection(cfg *amqp.Config, conn ConnectionConfig) {
	// 配置底层 AMQP 配置（仅在需要时）
	ConfigureAMQP(cfg, conn.AMQP)

	// 配置 TLS（仅在启用时）
	if conn.TLS.Enabled {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: conn.TLS.InsecureSkipVerify,
			ServerName:         conn.TLS.ServerName,
		}

		// 加载客户端证书和私钥（如果提供）
		if conn.TLS.CertFile != "" && conn.TLS.KeyFile != "" {
			cert, err := tls.LoadX509KeyPair(conn.TLS.CertFile, conn.TLS.KeyFile)
			if err == nil {
				tlsConfig.Certificates = []tls.Certificate{cert}
			}
		}

		// 注意：CA 证书加载需要额外的处理，这里简化处理
		// 如果需要完整的 CA 证书支持，可以使用 x509.CertPool

		cfg.Connection.TLSConfig = tlsConfig
	}
	// 如果 TLS.Enabled 为 false，使用 watermill-amqp 的默认值（nil，即不使用 TLS）

	// 配置重连（仅在启用时）
	if conn.Reconnect.Enabled {
		reconnectCfg := &amqp.ReconnectConfig{}

		// 解析时间间隔
		if conn.Reconnect.InitialInterval != "" {
			if d, err := time.ParseDuration(conn.Reconnect.InitialInterval); err == nil {
				reconnectCfg.BackoffInitialInterval = d
			} else {
				reconnectCfg.BackoffInitialInterval = 1 * time.Second // 默认值
			}
		} else {
			reconnectCfg.BackoffInitialInterval = 1 * time.Second
		}

		if conn.Reconnect.MaxInterval != "" {
			if d, err := time.ParseDuration(conn.Reconnect.MaxInterval); err == nil {
				reconnectCfg.BackoffMaxInterval = d
			} else {
				reconnectCfg.BackoffMaxInterval = 30 * time.Second // 默认值
			}
		} else {
			reconnectCfg.BackoffMaxInterval = 30 * time.Second
		}

		// 设置退避参数
		if conn.Reconnect.RandomizationFactor > 0 {
			reconnectCfg.BackoffRandomizationFactor = conn.Reconnect.RandomizationFactor
		} else {
			reconnectCfg.BackoffRandomizationFactor = 0.5 // 默认值
		}

		if conn.Reconnect.Multiplier > 0 {
			reconnectCfg.BackoffMultiplier = conn.Reconnect.Multiplier
		} else {
			reconnectCfg.BackoffMultiplier = 1.5 // 默认值
		}

		cfg.Connection.Reconnect = reconnectCfg
	}
	// 如果 Reconnect.Enabled 为 false，使用 watermill-amqp 的默认值（nil，即不自动重连）
}

// ConfigureAMQP 配置底层 AMQP 配置（amqp091-go 的 Config）。
// 仅在配置项有值时应用，否则使用 amqp091-go 的默认值。
func ConfigureAMQP(cfg *amqp.Config, amqpCfg AMQPConfig) {
	// 检查是否有任何配置需要设置
	hasConfig := amqpCfg.Vhost != "" || amqpCfg.Heartbeat > 0 || amqpCfg.Locale != "" ||
		amqpCfg.ChannelMax > 0 || amqpCfg.FrameSize > 0

	if !hasConfig {
		// 没有任何配置，使用默认值（amqp091-go 会使用其默认值）
		return
	}

	// 如果 Connection.AmqpConfig 为 nil，创建一个新的
	if cfg.Connection.AmqpConfig == nil {
		cfg.Connection.AmqpConfig = &amqp091.Config{}
	}

	amqpConfig := cfg.Connection.AmqpConfig

	// 配置 Vhost（仅在指定时）
	if amqpCfg.Vhost != "" {
		amqpConfig.Vhost = amqpCfg.Vhost
	}

	// 配置 Heartbeat（仅在指定时）
	if amqpCfg.Heartbeat > 0 {
		amqpConfig.Heartbeat = time.Duration(amqpCfg.Heartbeat) * time.Second
	}
	// 如果 Heartbeat 为 0 或未设置，使用 amqp091-go 的默认值

	// 配置 Locale（仅在指定时）
	if amqpCfg.Locale != "" {
		amqpConfig.Locale = amqpCfg.Locale
	}
	// 如果 Locale 为空，使用 amqp091-go 的默认值

	// 配置 ChannelMax（仅在指定时）
	if amqpCfg.ChannelMax > 0 {
		amqpConfig.ChannelMax = uint16(amqpCfg.ChannelMax)
	}
	// 如果 ChannelMax 为 0，使用 amqp091-go 的默认值（0 表示无限制）

	// 配置 FrameSize（仅在指定时）
	if amqpCfg.FrameSize > 0 {
		amqpConfig.FrameSize = amqpCfg.FrameSize
	}
	// 如果 FrameSize 为 0，使用 amqp091-go 的默认值
}

// ConfigurePublisher 配置发布者相关选项。
func ConfigurePublisher(cfg *amqp.Config, p ProducerConfig) {
	// 设置 Mandatory 和 Immediate 标志
	cfg.Publish.Mandatory = p.Mandatory
	cfg.Publish.Immediate = p.Immediate

	// 配置发布确认（RabbitMQ 特有功能）
	cfg.Publish.ConfirmDelivery = p.ConfirmDelivery

	// 配置通道池大小（RabbitMQ 特有功能，用于提高性能）
	if p.ChannelPoolSize > 0 {
		cfg.Publish.ChannelPoolSize = p.ChannelPoolSize
	}

	// 配置事务模式（RabbitMQ 特有功能）
	cfg.Publish.Transactional = p.Transactional

	// 注意：watermill-amqp 的 PublishConfig 没有直接的 Persistent 字段
	// 消息持久化是通过队列的 Durable 配置和消息的 DeliveryMode 控制的
	// DeliveryMode 在发布消息时通过 amqp091-go 库设置，这里我们只配置发布选项
}

// ConfigureConsumer 配置消费者相关选项。
func ConfigureConsumer(cfg *amqp.Config, c ConsumerConfig) {
	if c.QueueName != "" {
		cfg.Queue.GenerateName = func(topic string) string {
			return c.QueueName
		}
	}

	if c.ConsumerTag != "" {
		cfg.Consume.Consumer = c.ConsumerTag
	}

	cfg.Consume.Exclusive = c.Exclusive
	cfg.Consume.NoLocal = c.NoLocal
	cfg.Consume.NoWait = c.NoWait
	cfg.Consume.NoRequeueOnNack = c.NoRequeueOnNack

	if c.PrefetchCount > 0 {
		cfg.Consume.Qos.PrefetchCount = c.PrefetchCount
	}

	if c.PrefetchSize > 0 {
		cfg.Consume.Qos.PrefetchSize = c.PrefetchSize
	}

	// 配置 QOS Global 标志（RabbitMQ 特有功能）
	cfg.Consume.Qos.Global = c.QosGlobal

	// AutoAck 在 watermill-amqp 中通过 QOS 控制
	// 如果 AutoAck 为 true，则不设置 QOS（或者设置为不等待确认的模式）
	// 注意：watermill 通常管理 ack，所以这里主要设置预取数量
}

// ConfigureExchange 配置交换机相关选项。
func ConfigureExchange(cfg *amqp.Config, e ExchangeConfig) {
	exchangeType := e.Type
	if exchangeType == "" {
		exchangeType = messaging.DefaultExchangeType // 默认使用 topic 交换机
	}

	cfg.Exchange.Type = exchangeType.String()

	// 如果指定了交换机名称，使用 GenerateName 函数返回固定名称
	if e.Name != "" {
		cfg.Exchange.GenerateName = func(topic string) string {
			return e.Name
		}
	}

	// 设置交换机持久化（默认 true）
	cfg.Exchange.Durable = e.Durable
	if !e.Durable {
		cfg.Exchange.Durable = true // 默认持久化
	}

	cfg.Exchange.AutoDeleted = e.AutoDelete
	cfg.Exchange.Internal = e.Internal
	cfg.Exchange.NoWait = e.NoWait
}

// ConfigureQueue 配置队列相关选项（RabbitMQ 特有）。
func ConfigureQueue(cfg *amqp.Config, q QueueConfig, c ConsumerConfig) {
	// 队列名称在 ConfigureConsumer 中已设置，这里只配置队列属性

	// 设置队列持久化（默认 true，与 NewDurableQueueConfig 保持一致）
	cfg.Queue.Durable = q.Durable
	if !q.Durable {
		cfg.Queue.Durable = true // 默认持久化
	}

	cfg.Queue.AutoDelete = q.AutoDelete
	cfg.Queue.Exclusive = q.Exclusive
	cfg.Queue.NoWait = q.NoWait

	// 配置队列 Arguments（RabbitMQ 高级特性）
	hasArguments := q.MaxPriority > 0 || q.MessageTTL > 0 || q.MaxLength > 0 ||
		q.MaxLengthBytes > 0 || q.DeadLetterExchange != "" || q.DeadLetterRoutingKey != ""

	if hasArguments {
		if cfg.Queue.Arguments == nil {
			cfg.Queue.Arguments = make(amqp091.Table)
		}

		// 优先级队列（Priority Queue）
		if q.MaxPriority > 0 {
			cfg.Queue.Arguments["x-max-priority"] = int32(q.MaxPriority)
		}

		// 延迟队列（Message TTL）
		if q.MessageTTL > 0 {
			cfg.Queue.Arguments["x-message-ttl"] = int32(q.MessageTTL)
		}

		// 队列长度限制
		if q.MaxLength > 0 {
			cfg.Queue.Arguments["x-max-length"] = int32(q.MaxLength)
		}

		if q.MaxLengthBytes > 0 {
			cfg.Queue.Arguments["x-max-length-bytes"] = int32(q.MaxLengthBytes)
		}

		// 死信队列（Dead Letter Queue）
		if q.DeadLetterExchange != "" {
			cfg.Queue.Arguments["x-dead-letter-exchange"] = q.DeadLetterExchange
		}

		if q.DeadLetterRoutingKey != "" {
			cfg.Queue.Arguments["x-dead-letter-routing-key"] = q.DeadLetterRoutingKey
		}
	}
}

// ConfigureQueueBind 配置队列绑定相关选项（RabbitMQ 特有）。
func ConfigureQueueBind(cfg *amqp.Config, qb QueueBindConfig) {
	// 配置路由键生成函数
	if qb.RoutingKey != "" {
		cfg.QueueBind.GenerateRoutingKey = func(topic string) string {
			return qb.RoutingKey
		}
	}
	// 如果 RoutingKey 为空，使用默认行为（使用 topic 作为 routing key）

	cfg.QueueBind.NoWait = qb.NoWait

	// 注意：Arguments 可以通过 amqp.Table 设置，用于：
	// - 自定义路由匹配规则
	// - Headers exchange 的匹配条件
	// 如果需要这些功能，可以在 QueueBindConfig 中添加相应字段
}
