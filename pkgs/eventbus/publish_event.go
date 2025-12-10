package eventbus

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

const DefaultEventTypeHeaderKey = "event_type"
const DefaultPartitionKey = "partition"

type PublishOption func(*message.Message)

func WithMetadata(meta map[string]string) PublishOption {
	return func(m *message.Message) {
		for k, v := range meta {
			m.Metadata.Set(k, v)
		}
	}
}

func WithMessageID(id string) PublishOption {
	return func(m *message.Message) { m.UUID = id }
}

func WithPartitionKey(pk string) PublishOption {
	return func(m *message.Message) { m.Metadata.Set("partition_key", pk) }
}

func WithEventTypeHeaderKey(key string) PublishOption {
	return func(m *message.Message) { m.Metadata.Set("__event_type_header_key__", key) }
}

func PublishEvent[T TypedEvent](
	ctx context.Context,
	bp *BusPublisher,
	evt T,
	opts ...PublishOption,
) error {
	topicName, ok := bp.GetTopic(TopicKeyOf[T]())
	if !ok || topicName == "" {
		return fmt.Errorf("eventbus: topic not found for event type %T", evt)
	}

	payload, err := json.Marshal(evt)
	if err != nil {
		return fmt.Errorf("eventbus: marshal failed: %w", err)
	}

	id, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("eventbus: generate id failed: %w", err)
	}

	msg := message.NewMessageWithContext(ctx, id.String(), payload)
	for _, opt := range opts {
		opt(msg)
	}

	headerKey := msg.Metadata.Get("__event_type_header_key__")
	if headerKey == "" {
		headerKey = DefaultEventTypeHeaderKey
	}

	if msg.Metadata.Get(headerKey) == "" {
		msg.Metadata.Set(headerKey, EventTypeOf[T]())
	}

	return bp.Publish(topicName, msg)
}
