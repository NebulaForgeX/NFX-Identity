package events

import (
	"context"
	"time"
	eventCommands "nfxid/modules/audit/application/events/commands"
	eventDomain "nfxid/modules/audit/domain/events"

	"github.com/google/uuid"
)

// CreateEvent 创建事件
func (s *Service) CreateEvent(ctx context.Context, cmd eventCommands.CreateEventCmd) (uuid.UUID, error) {
	// Parse occurred at
	var occurredAt time.Time
	if cmd.OccurredAt != "" {
		parsed, err := time.Parse(time.RFC3339, cmd.OccurredAt)
		if err != nil {
			return uuid.Nil, err
		}
		occurredAt = parsed
	}

	// Create domain entity
	event, err := eventDomain.NewEvent(eventDomain.NewEventParams{
		EventID:              cmd.EventID,
		OccurredAt:           occurredAt,
		TenantID:             cmd.TenantID,
		AppID:                cmd.AppID,
		ActorType:            cmd.ActorType,
		ActorID:              cmd.ActorID,
		ActorTenantMemberID:  cmd.ActorTenantMemberID,
		Action:               cmd.Action,
		TargetType:           cmd.TargetType,
		TargetID:             cmd.TargetID,
		Result:               cmd.Result,
		FailureReasonCode:    cmd.FailureReasonCode,
		HTTPMethod:           cmd.HTTPMethod,
		HTTPPath:             cmd.HTTPPath,
		HTTPStatus:           cmd.HTTPStatus,
		RequestID:            cmd.RequestID,
		TraceID:              cmd.TraceID,
		IP:                   cmd.IP,
		UserAgent:            cmd.UserAgent,
		GeoCountry:           cmd.GeoCountry,
		RiskLevel:            cmd.RiskLevel,
		DataClassification:   cmd.DataClassification,
		PrevHash:             cmd.PrevHash,
		EventHash:            cmd.EventHash,
		Metadata:             cmd.Metadata,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.eventRepo.Create.New(ctx, event); err != nil {
		return uuid.Nil, err
	}

	return event.ID(), nil
}
