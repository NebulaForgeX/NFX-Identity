package eventbus

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

type BusSubscriber struct {
	sub           message.Subscriber
	topicResolver *TopicResolver
}

func NewSubscriber(
	sub message.Subscriber,
	topicResolver *TopicResolver,
) *BusSubscriber {
	return &BusSubscriber{sub, topicResolver}
}

func (s *BusSubscriber) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	return s.sub.Subscribe(ctx, topic)
}

func (s *BusSubscriber) Close() error {
	return s.sub.Close()
}

func (s *BusSubscriber) GetTopic(key TopicKey) (string, bool) {
	return s.topicResolver.GetName(key)
}
