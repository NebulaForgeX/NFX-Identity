package handler

import (
	"context"
	"time"

	eventApp "nfxid/modules/audit/application/events"
	"nfxid/modules/audit/interfaces/grpc/mapper"
	eventpb "nfxid/protos/gen/audit/event"
	"nfxid/pkgs/errx"

	"github.com/google/uuid"
)

type EventHandler struct {
	eventpb.UnimplementedEventServiceServer
	eventAppSvc *eventApp.Service
}

func NewEventHandler(eventAppSvc *eventApp.Service) *EventHandler {
	return &EventHandler{
		eventAppSvc: eventAppSvc,
	}
}

// GetEventByID 根据ID获取事件
func (h *EventHandler) GetEventByID(ctx context.Context, req *eventpb.GetEventByIDRequest) (*eventpb.GetEventByIDResponse, error) {
	eventID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	eventView, err := h.eventAppSvc.GetEvent(ctx, eventID)
	if err != nil {
		return nil, err
	}

	event := mapper.EventROToProto(&eventView)
	return &eventpb.GetEventByIDResponse{Event: event}, nil
}

// GetEventByEventID 根据事件ID获取事件
func (h *EventHandler) GetEventByEventID(ctx context.Context, req *eventpb.GetEventByEventIDRequest) (*eventpb.GetEventByEventIDResponse, error) {
	eventView, err := h.eventAppSvc.GetEventByEventID(ctx, req.EventId)
	if err != nil {
		return nil, err
	}

	event := mapper.EventROToProto(&eventView)
	return &eventpb.GetEventByEventIDResponse{Event: event}, nil
}

// GetEventsByActor 根据操作者获取事件列表
func (h *EventHandler) GetEventsByActor(ctx context.Context, req *eventpb.GetEventsByActorRequest) (*eventpb.GetEventsByActorResponse, error) {
	actorID, err := uuid.Parse(req.ActorId)
	if err != nil {
		return nil, errx.ErrInvalidParams.WithCause(err)
	}

	// Parse actor type
	actorType := mapper.ActorTypeFromProto(req.ActorType)

	// ByActor method doesn't require time range, pass nil
	var startTime, endTime *time.Time

	eventViews, err := h.eventAppSvc.GetEventsByActor(ctx, actorType, actorID, startTime, endTime)
	if err != nil {
		return nil, err
	}

	// Apply limit if specified
	if req.Limit != nil && *req.Limit > 0 && int(*req.Limit) < len(eventViews) {
		eventViews = eventViews[:*req.Limit]
	}

	events := mapper.EventListROToProto(eventViews)
	return &eventpb.GetEventsByActorResponse{Events: events}, nil
}

// BatchGetEvents 批量获取事件
func (h *EventHandler) BatchGetEvents(ctx context.Context, req *eventpb.BatchGetEventsRequest) (*eventpb.BatchGetEventsResponse, error) {
	eventIDs := make([]uuid.UUID, 0, len(req.Ids))
	for _, idStr := range req.Ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue
		}
		eventIDs = append(eventIDs, id)
	}

	events := make([]*eventpb.Event, 0, len(eventIDs))
	for _, eventID := range eventIDs {
		eventView, err := h.eventAppSvc.GetEvent(ctx, eventID)
		if err != nil {
			continue
		}
		event := mapper.EventROToProto(&eventView)
		events = append(events, event)
	}

	return &eventpb.BatchGetEventsResponse{Events: events}, nil
}
