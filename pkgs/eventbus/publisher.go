package eventbus

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type BusPublisher struct {
	pub           message.Publisher
	topicResolver *TopicResolver
}

func NewBusPublisher(
	pub message.Publisher,
	topicResolver *TopicResolver,
) *BusPublisher {
	return &BusPublisher{pub, topicResolver}
}

func (p *BusPublisher) Publish(topic string, msgs ...*message.Message) error {
	return p.pub.Publish(topic, msgs...)
}

func (p *BusPublisher) Close() error {
	return p.pub.Close()
}

func (p *BusPublisher) GetTopic(key TopicKey) (string, bool) {
	return p.topicResolver.GetName(key)
}
