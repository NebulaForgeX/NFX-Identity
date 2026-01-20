package rabbitmqx

import (
	"nfxid/pkgs/logx"
	"nfxid/pkgs/rabbitmqx/messaging"

	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
)

func NewSubscriber(cfg *Config) (*messaging.BusSubscriber, error) {
	uri, err := cfg.BuildURI()
	if err != nil {
		logx.S().Errorf("❌ Failed to build RabbitMQ URI: %v", err)
		return nil, err
	}

	amqpConfig, err := BuildAMQPConfig(cfg)
	if err != nil {
		logx.S().Errorf("❌ Failed to build AMQP config: %v", err)
		return nil, err
	}

	defaultQueue := cfg.Consumer.QueueName
	exchangeName := cfg.Exchange.Name
	if exchangeName == "" {
		exchangeName = "default"
	}
	logx.S().Infof("🔄 Initializing RabbitMQ Subscriber with URI: %s (default queue: %s, exchange: %s)", maskURI(uri), defaultQueue, exchangeName)

	// 使用支持优先级的 Marshaler（用于反序列化时保留优先级信息）
	amqpConfig.Marshaler = &PriorityMarshaler{}

	sub, err := amqp.NewSubscriber(amqpConfig, logx.NewZapWatermillLogger(logx.L()))
	if err != nil {
		logx.S().Errorf("❌ Failed to create RabbitMQ Subscriber: %v", err)
		return nil, err
	}

	// 将 ConsumerQueues 转换为 QueueResolver 需要的格式
	keyToBinding := make(map[messaging.MessageKey]messaging.QueueBinding)
	for eventKey, binding := range cfg.ConsumerQueues {
		queue := binding.Queue
		if queue == "" {
			queue = defaultQueue
			if queue == "" {
				queue = eventKey // 默认使用事件键作为 Queue
			}
		}
		bindingKey := binding.BindingKey
		if bindingKey == "" {
			bindingKey = eventKey // 默认使用事件键作为 BindingKey
		}
		keyToBinding[messaging.MessageKey(eventKey)] = messaging.QueueBinding{
			Queue:      queue,
			BindingKey: bindingKey,
		}
	}

	queueResolver, err := messaging.NewQueueResolver(keyToBinding)
	if err != nil {
		logx.S().Errorf("❌ Failed to create queue resolver: %v", err)
		return nil, err
	}

	logx.S().Infof("✅ Successfully connected to RabbitMQ Subscriber: %s (queue: %s, exchange: %s)", maskURI(uri), defaultQueue, exchangeName)
	return messaging.NewSubscriber(sub, queueResolver), nil
}
