package messaging

import (
	"context"

	"nfxid/pkgs/rabbitmqx/messaging"
)

type Service struct {
	publisher *messaging.BusPublisher
}

func NewService(publisher *messaging.BusPublisher) *Service {
	return &Service{
		publisher: publisher,
	}
}

// PublishMessage 发布消息到消息总线（泛型包装函数）
func PublishMessage[T messaging.TypedMessage](ctx context.Context, svc *Service, msg T, opts ...messaging.PublishOption) error {
	return messaging.PublishMessage(ctx, svc.publisher, msg, opts...)
}
