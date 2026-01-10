package messaging

// ExchangeType Exchange 类型
// 支持 RabbitMQ 的所有 Exchange 类型，包括基本类型和插件类型
type ExchangeType string

// 基本 Exchange 类型（RabbitMQ 内置）
const (
	// ExchangeTypeDirect Direct Exchange（直连交换机）
	// 精确匹配 RoutingKey，完全匹配才路由
	ExchangeTypeDirect ExchangeType = "direct"

	// ExchangeTypeTopic Topic Exchange（主题交换机）- 默认类型
	// 支持通配符匹配（* 和 #），最灵活的路由方式
	ExchangeTypeTopic ExchangeType = "topic"

	// ExchangeTypeFanout Fanout Exchange（扇出交换机）
	// 忽略 RoutingKey，广播到所有绑定的队列
	ExchangeTypeFanout ExchangeType = "fanout"

	// ExchangeTypeHeaders Headers Exchange（头交换机）
	// 根据消息头（Headers）匹配，忽略 RoutingKey
	ExchangeTypeHeaders ExchangeType = "headers"
)

// 插件 Exchange 类型（需要安装对应的 RabbitMQ 插件）
const (
	// ExchangeTypeDelayedMessage 延迟消息 Exchange（需要 rabbitmq-delayed-message-exchange 插件）
	// 支持延迟消息投递，通过 x-delay 消息头指定延迟时间（毫秒）
	// 安装插件：rabbitmq-plugins enable rabbitmq_delayed_message_exchange
	ExchangeTypeDelayedMessage ExchangeType = "x-delayed-message"

	// ExchangeTypeConsistentHash 一致性哈希 Exchange（需要 rabbitmq-consistent-hash-exchange 插件）
	// 根据 RoutingKey 的哈希值路由消息，实现负载均衡
	// 安装插件：rabbitmq-plugins enable rabbitmq_consistent_hash_exchange
	ExchangeTypeConsistentHash ExchangeType = "x-consistent-hash"

	// ExchangeTypeSharding 分片 Exchange（需要 rabbitmq-sharding 插件）
	// 将消息分片到多个队列，提高并发处理能力
	// 安装插件：rabbitmq-plugins enable rabbitmq_sharding
	ExchangeTypeSharding ExchangeType = "x-sharding"

	// ExchangeTypeModulusHash 模数哈希 Exchange（需要 rabbitmq-modulus-hash-exchange 插件）
	// 根据 RoutingKey 的模数路由消息
	ExchangeTypeModulusHash ExchangeType = "x-modulus-hash"

	// ExchangeTypeRandom 随机 Exchange（需要 rabbitmq-random-exchange 插件）
	// 随机路由消息到绑定的队列
	ExchangeTypeRandom ExchangeType = "x-random"

	// ExchangeTypeJMS 主题 Exchange（需要 rabbitmq-jms-topic-exchange 插件）
	// 支持 JMS 主题语义
	ExchangeTypeJMS ExchangeType = "x-jms-topic"
)

// String 返回 Exchange 类型的字符串表示
func (et ExchangeType) String() string {
	return string(et)
}

// IsValid 检查 Exchange 类型是否有效
// 注意：插件类型需要安装对应的插件才能使用
func (et ExchangeType) IsValid() bool {
	validTypes := map[ExchangeType]bool{
		// 基本类型
		ExchangeTypeDirect:  true,
		ExchangeTypeTopic:   true,
		ExchangeTypeFanout:  true,
		ExchangeTypeHeaders: true,
		// 插件类型
		ExchangeTypeDelayedMessage: true,
		ExchangeTypeConsistentHash: true,
		ExchangeTypeSharding:       true,
		ExchangeTypeModulusHash:    true,
		ExchangeTypeRandom:         true,
		ExchangeTypeJMS:            true,
	}
	return validTypes[et]
}

// IsPluginType 检查是否为插件类型
func (et ExchangeType) IsPluginType() bool {
	pluginTypes := map[ExchangeType]bool{
		ExchangeTypeDelayedMessage: true,
		ExchangeTypeConsistentHash: true,
		ExchangeTypeSharding:       true,
		ExchangeTypeModulusHash:    true,
		ExchangeTypeRandom:         true,
		ExchangeTypeJMS:            true,
	}
	return pluginTypes[et]
}

// IsBasicType 检查是否为基本类型
func (et ExchangeType) IsBasicType() bool {
	basicTypes := map[ExchangeType]bool{
		ExchangeTypeDirect:  true,
		ExchangeTypeTopic:   true,
		ExchangeTypeFanout:  true,
		ExchangeTypeHeaders: true,
	}
	return basicTypes[et]
}

// DefaultExchangeType 默认 Exchange 类型
const DefaultExchangeType = ExchangeTypeTopic
