package rabbitmqx

import "nfxid/pkgs/rabbitmqx/messaging"

type Config struct {
	URI        string           `koanf:"uri"`
	ClientID   string           `koanf:"client_id"`
	Producer   ProducerConfig   `koanf:"producer"`
	Consumer   ConsumerConfig   `koanf:"consumer"`
	Connection ConnectionConfig `koanf:"connection"`
	Exchange   ExchangeConfig   `koanf:"exchange"`
	Queue      QueueConfig      `koanf:"queue"`
	QueueBind  QueueBindConfig  `koanf:"queue_bind"`
	// ProducerExchanges 映射事件键到 Exchange 和 RoutingKey
	// 格式: event_key -> exchange:routing_key 或 routing_key (如果 exchange 为空则使用默认)
	ProducerExchanges map[messaging.MessageKey]ProducerRouting `koanf:"producer_exchanges"`
	// ConsumerQueues 映射事件键到 Queue 和 BindingKey
	// 格式: event_key -> queue:binding_key 或 binding_key (如果 queue 为空则自动生成)
	ConsumerQueues map[messaging.MessageKey]ConsumerBinding `koanf:"consumer_queues"`
}

// ProducerRouting 定义发布者的路由配置（Exchange 和 RoutingKey）
type ProducerRouting struct {
	Exchange   string `koanf:"exchange"`    // Exchange 名称，为空则使用 ExchangeConfig.Name 或根据事件键生成
	RoutingKey string `koanf:"routing_key"` // RoutingKey，为空则使用事件键作为 RoutingKey
}

// ConsumerBinding 定义消费者的绑定配置（Queue 和 BindingKey）
type ConsumerBinding struct {
	Queue      string `koanf:"queue"`       // Queue 名称，为空则根据事件键自动生成
	BindingKey string `koanf:"binding_key"` // BindingKey（用于匹配 RoutingKey），为空则使用事件键
}

type ProducerConfig struct {
	Mandatory       bool   `koanf:"mandatory"`         // 如果为 true，消息无法路由时会返回错误
	Immediate       bool   `koanf:"immediate"`         // 如果为 true，消息无法立即投递给消费者时会返回错误
	ContentType     string `koanf:"content_type"`      // 消息内容类型，默认 "application/json"（注意：watermill-amqp 可能不直接支持，需通过 Marshaler 设置）
	DeliveryMode    uint8  `koanf:"delivery_mode"`     // 1=非持久化, 2=持久化，默认 2（注意：通过队列 Durable 和消息属性控制）
	ConfirmDelivery bool   `koanf:"confirm_delivery"`  // 是否等待服务器确认（对应 PublishConfig.ConfirmDelivery）
	ChannelPoolSize int    `koanf:"channel_pool_size"` // 通道池大小，0 表示不池化（对应 PublishConfig.ChannelPoolSize）
	Transactional   bool   `koanf:"transactional"`     // 是否启用事务模式（对应 PublishConfig.Transactional）
	DefaultPriority uint8  `koanf:"default_priority"`  // 默认消息优先级（0-255），0 表示不设置优先级（RabbitMQ 特有，需通过消息元数据设置）
}

type ConsumerConfig struct {
	QueueName       string `koanf:"queue_name"`         // 队列名称（用于消费者），如果为空则根据 topic 生成
	ConsumerTag     string `koanf:"consumer_tag"`       // 消费者标签（对应 ConsumeConfig.Consumer）
	AutoAck         bool   `koanf:"auto_ack"`           // 是否自动确认，默认 false（注意：watermill 通常管理 ack）
	Exclusive       bool   `koanf:"exclusive"`          // 是否独占队列
	NoLocal         bool   `koanf:"no_local"`           // 不接收同连接发布的消息
	NoWait          bool   `koanf:"no_wait"`            // 不等待服务器响应
	NoRequeueOnNack bool   `koanf:"no_requeue_on_nack"` // 当 nack 时不重新入队（对应 ConsumeConfig.NoRequeueOnNack）
	PrefetchCount   int    `koanf:"prefetch_count"`     // 预取数量，0 表示不限制（对应 QosConfig.PrefetchCount）
	PrefetchSize    int    `koanf:"prefetch_size"`      // 预取大小（字节），0 表示不限制（对应 QosConfig.PrefetchSize）
	QosGlobal       bool   `koanf:"qos_global"`         // QOS 是否全局应用（对应 QosConfig.Global）
}

type ConnectionConfig struct {
	// TLS 配置
	TLS TLSConfig `koanf:"tls"`

	// 重连配置
	Reconnect ReconnectConfig `koanf:"reconnect"`

	// 底层 AMQP 配置（amqp091-go 的 Config）
	AMQP AMQPConfig `koanf:"amqp"`

	// 注意：AmqpURI 在 Config.URI 中设置
	// Vhost 可以在 URI 中指定，例如: amqp://user:pass@host:port/vhost
	// 或者通过 AMQP.Vhost 配置
}

type AMQPConfig struct {
	Vhost      string `koanf:"vhost"`       // 虚拟主机，默认 "/"（也可以通过 URI 指定）
	Heartbeat  int    `koanf:"heartbeat"`   // 心跳间隔（秒），默认 10，0 表示禁用
	Locale     string `koanf:"locale"`      // 区域设置，默认 "en_US"
	ChannelMax int    `koanf:"channel_max"` // 最大通道数，0 表示无限制
	FrameSize  int    `koanf:"frame_size"`  // 最大帧大小，0 表示使用默认值
}

type TLSConfig struct {
	Enabled            bool   `koanf:"enabled"`              // 是否启用 TLS
	InsecureSkipVerify bool   `koanf:"insecure_skip_verify"` // 是否跳过证书验证（仅用于开发环境）
	CertFile           string `koanf:"cert_file"`            // 客户端证书文件路径
	KeyFile            string `koanf:"key_file"`             // 客户端私钥文件路径
	CAFile             string `koanf:"ca_file"`              // CA 证书文件路径
	ServerName         string `koanf:"server_name"`          // 服务器名称（用于证书验证）
}

type ReconnectConfig struct {
	Enabled             bool    `koanf:"enabled"`              // 是否启用自动重连
	InitialInterval     string  `koanf:"initial_interval"`     // 初始重连间隔，例如 "1s"
	RandomizationFactor float64 `koanf:"randomization_factor"` // 随机化因子，默认 0.5
	Multiplier          float64 `koanf:"multiplier"`           // 退避乘数，默认 1.5
	MaxInterval         string  `koanf:"max_interval"`         // 最大重连间隔，例如 "30s"
}

type ExchangeConfig struct {
	Name       string `koanf:"name"`        // 交换机名称，默认 ""（如果为空则根据 topic 生成，对应 ExchangeConfig.GenerateName）
	Type       string `koanf:"type"`        // 交换机类型: direct, topic, fanout, headers，默认 "topic"
	Durable    bool   `koanf:"durable"`     // 是否持久化，默认 true（对应 ExchangeConfig.Durable）
	AutoDelete bool   `koanf:"auto_delete"` // 是否自动删除，默认 false（对应 ExchangeConfig.AutoDeleted）
	Internal   bool   `koanf:"internal"`    // 是否内部交换机，默认 false（对应 ExchangeConfig.Internal）
	NoWait     bool   `koanf:"no_wait"`     // 不等待服务器响应，默认 false（对应 ExchangeConfig.NoWait）
	// Arguments 可以通过 amqp.Table 设置，这里简化处理
}

type QueueConfig struct {
	// 注意：QueueName 在 ConsumerConfig.QueueName 中设置
	Durable    bool `koanf:"durable"`     // 是否持久化，默认 true（对应 QueueConfig.Durable）
	AutoDelete bool `koanf:"auto_delete"` // 是否自动删除，默认 false（对应 QueueConfig.AutoDelete）
	Exclusive  bool `koanf:"exclusive"`   // 是否独占队列，默认 false（对应 QueueConfig.Exclusive）
	NoWait     bool `koanf:"no_wait"`     // 不等待服务器响应，默认 false（对应 QueueConfig.NoWait）

	// 优先级队列（Priority Queue）
	MaxPriority int `koanf:"max_priority"` // 队列最大优先级（0-255），0 表示不启用优先级队列（对应 x-max-priority）

	// 延迟队列（Delayed Queue / Message TTL）
	MessageTTL int `koanf:"message_ttl"` // 消息 TTL（毫秒），0 表示不设置 TTL（对应 x-message-ttl）

	// 队列限制
	MaxLength      int `koanf:"max_length"`       // 队列最大消息数，0 表示不限制（对应 x-max-length）
	MaxLengthBytes int `koanf:"max_length_bytes"` // 队列最大字节数，0 表示不限制（对应 x-max-length-bytes）

	// 死信队列（Dead Letter Queue）
	DeadLetterExchange   string `koanf:"dead_letter_exchange"`    // 死信交换机名称（对应 x-dead-letter-exchange）
	DeadLetterRoutingKey string `koanf:"dead_letter_routing_key"` // 死信路由键（对应 x-dead-letter-routing-key）

	// 延迟消息插件支持（RabbitMQ Delayed Message Plugin）
	// 注意：需要安装 rabbitmq-delayed-message-exchange 插件
	// 如果使用延迟插件，应该通过 Exchange 类型 "x-delayed-message" 实现，而不是队列 TTL
}

type QueueBindConfig struct {
	RoutingKey string `koanf:"routing_key"` // 路由键，如果为空则使用 topic 名称（对应 QueueBindConfig.GenerateRoutingKey）
	NoWait     bool   `koanf:"no_wait"`     // 不等待服务器响应，默认 false（对应 QueueBindConfig.NoWait）
	// Arguments 可以通过 amqp.Table 设置，用于路由匹配规则
}
