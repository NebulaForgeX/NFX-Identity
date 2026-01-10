package rabbitmqx

import (
	"strconv"

	"nfxid/pkgs/rabbitmqx/messaging"

	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

// PriorityMarshaler 支持优先级的消息序列化器
// 从消息元数据中读取优先级并设置到 AMQP 消息属性
type PriorityMarshaler struct {
	amqp.DefaultMarshaler
}

// Marshal 序列化消息并设置优先级
func (m *PriorityMarshaler) Marshal(msg *message.Message) (amqp091.Publishing, error) {
	pub, err := m.DefaultMarshaler.Marshal(msg)
	if err != nil {
		return amqp091.Publishing{}, err
	}

	// 从消息元数据中读取优先级
	if priorityStr := msg.Metadata.Get(messaging.MessagePriorityKey); priorityStr != "" {
		if priority, err := strconv.ParseUint(priorityStr, 10, 8); err == nil {
			pub.Priority = uint8(priority)
		}
	}

	return pub, nil
}

// Unmarshal 反序列化消息
func (m *PriorityMarshaler) Unmarshal(amqpMsg amqp091.Delivery) (*message.Message, error) {
	msg, err := m.DefaultMarshaler.Unmarshal(amqpMsg)
	if err != nil {
		return nil, err
	}

	// 将优先级写回消息元数据（如果需要）
	if amqpMsg.Priority > 0 {
		msg.Metadata.Set(messaging.MessagePriorityKey, strconv.Itoa(int(amqpMsg.Priority)))
	}

	return msg, nil
}
