package eventbus

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

type EventType = string

type EventHandler[T TypedEvent] func(ctx context.Context, evt T, msg *message.Message) error

type Validatable interface {
	Validate() error
}

type TypedEvent interface {
	EventType() EventType
	TopicKey() TopicKey
}

func EventTypeOf[T TypedEvent]() string {
	var zero T
	return zero.EventType()
}

func TopicKeyOf[T TypedEvent]() TopicKey {
	var zero T
	return zero.TopicKey()
}
