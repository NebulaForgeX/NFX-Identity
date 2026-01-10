package messaging

import "fmt"

// QueueResolver Queue 解析器。
// 用于在消息键（MessageKey）和 Queue + BindingKey 之间进行映射。
type QueueResolver struct {
	KeyToQueue map[MessageKey]QueueBinding // 消息键到 Queue + BindingKey 的映射
	QueueToKey map[string]MessageKey       // Queue+BindingKey 组合到消息键的映射（用于反向查找）
}

// QueueBinding 包含 Queue 名称和 BindingKey
type QueueBinding struct {
	Queue      string
	BindingKey string
}

// NewQueueResolver 创建新的 Queue 解析器。
// 根据提供的键值对映射创建双向映射。
//
// 参数:
//   - keyToBinding: 事件键到 QueueBinding 的映射
//
// 返回:
//   - *QueueResolver: Queue 解析器实例
//   - error: 如果 Queue 或 BindingKey 为空或存在重复，返回错误
//
// 示例:
//
//	resolver, err := messaging.NewQueueResolver(map[messaging.MessageKey]messaging.QueueBinding{
//	    "user_created": {Queue: "user-queue", BindingKey: "user.created"},
//	    "order_paid":   {Queue: "order-queue", BindingKey: "order.paid"},
//	})
func NewQueueResolver(keyToBinding map[MessageKey]QueueBinding) (*QueueResolver, error) {
	queueToKey := make(map[string]MessageKey)
	for k, binding := range keyToBinding {
		if binding.Queue == "" {
			return nil, fmt.Errorf("queue name is empty for key %q", k)
		}
		if binding.BindingKey == "" {
			return nil, fmt.Errorf("binding key is empty for key %q", k)
		}
		// 使用 Queue:BindingKey 作为唯一键
		key := binding.Queue + ":" + binding.BindingKey
		if _, dup := queueToKey[key]; dup {
			return nil, fmt.Errorf("duplicate queue:binding_key combination %q", key)
		}
		queueToKey[key] = k
	}
	return &QueueResolver{
		KeyToQueue: keyToBinding,
		QueueToKey: queueToKey,
	}, nil
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
//	binding, ok := resolver.GetQueue("user_created")
//	if ok {
//	    fmt.Println("Queue:", binding.Queue, "BindingKey:", binding.BindingKey)
//	}
func (q *QueueResolver) GetQueue(key MessageKey) (QueueBinding, bool) {
	v, ok := q.KeyToQueue[key]
	return v, ok
}

// GetKey 根据 Queue 和 BindingKey 获取消息键。
//
// 参数:
//   - queue: Queue 名称
//   - bindingKey: BindingKey
//
// 返回:
//   - MessageKey: 消息键，如果找到
//   - bool: 是否找到对应的消息键
func (q *QueueResolver) GetKey(queue, bindingKey string) (MessageKey, bool) {
	key := queue + ":" + bindingKey
	v, ok := q.QueueToKey[key]
	return v, ok
}
