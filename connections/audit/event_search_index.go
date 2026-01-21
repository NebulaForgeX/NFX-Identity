package audit

import (
	"context"
	"fmt"

	eventsearchindexpb "nfxid/protos/gen/audit/event_search_index"
)

// EventSearchIndexClient EventSearchIndex 客户端
type EventSearchIndexClient struct {
	client eventsearchindexpb.EventSearchIndexServiceClient
}

// NewEventSearchIndexClient 创建 EventSearchIndex 客户端
func NewEventSearchIndexClient(client eventsearchindexpb.EventSearchIndexServiceClient) *EventSearchIndexClient {
	return &EventSearchIndexClient{client: client}
}

// GetEventSearchIndexByID 根据ID获取索引
func (c *EventSearchIndexClient) GetEventSearchIndexByID(ctx context.Context, id string) (*eventsearchindexpb.EventSearchIndex, error) {
	req := &eventsearchindexpb.GetEventSearchIndexByIDRequest{
		Id: id,
	}

	resp, err := c.client.GetEventSearchIndexByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.EventSearchIndex, nil
}

// GetEventSearchIndexByEventID 根据事件ID获取索引
func (c *EventSearchIndexClient) GetEventSearchIndexByEventID(ctx context.Context, eventID string) (*eventsearchindexpb.EventSearchIndex, error) {
	req := &eventsearchindexpb.GetEventSearchIndexByEventIDRequest{
		EventId: eventID,
	}

	resp, err := c.client.GetEventSearchIndexByEventID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp.EventSearchIndex, nil
}

// SearchEvents 搜索事件
func (c *EventSearchIndexClient) SearchEvents(ctx context.Context, req *eventsearchindexpb.SearchEventsRequest) (*eventsearchindexpb.SearchEventsResponse, error) {
	resp, err := c.client.SearchEvents(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC call failed: %w", err)
	}

	return resp, nil
}