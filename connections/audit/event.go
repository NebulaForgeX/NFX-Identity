package audit

import (
	"context"
	"fmt"

	eventpb "nfxid/protos/gen/audit/event"
)

// EventClient Event 客户端
type EventClient struct {
	client eventpb.EventServiceClient
}

// NewEventClient 创建 Event 客户端
func NewEventClient(client eventpb.EventServiceClient) *EventClient {
	return &EventClient{client: client}
}

// GetEventByID 根据ID获取事件
func (c *EventClient) GetEventByID(ctx context.Context, id string) (*eventpb.Event, error) {
	req := &eventpb.GetEventByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetEventByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Event, nil
}

// GetEventByEventID 根据事件ID获取事件
func (c *EventClient) GetEventByEventID(ctx context.Context, eventID string) (*eventpb.Event, error) {
	req := &eventpb.GetEventByEventIDRequest{
		EventId: eventID,
	}

	resp, err := c.client.GetEventByEventID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Event, nil
}

// GetEventsByActor 根据操作者获取事件列表
func (c *EventClient) GetEventsByActor(ctx context.Context, actorType eventpb.AuditActorType, actorID string, tenantID *string, limit *int32) ([]*eventpb.Event, error) {
	req := &eventpb.GetEventsByActorRequest{
		ActorType: actorType,
		ActorId:   actorID,
		TenantId:  tenantID,
		Limit:     limit,
	}

	resp, err := c.client.GetEventsByActor(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Events, nil
}

// BatchGetEvents 批量获取事件
func (c *EventClient) BatchGetEvents(ctx context.Context, ids []string) ([]*eventpb.Event, error) {
	req := &eventpb.BatchGetEventsRequest{
		Ids: ids,
	}

	resp, err := c.client.BatchGetEvents(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.Events, nil
}