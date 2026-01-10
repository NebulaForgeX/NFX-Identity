package event_search_index

import (
	"context"
	"time"
	eventSearchIndexCommands "nfxid/modules/audit/application/event_search_index/commands"
	eventSearchIndexDomain "nfxid/modules/audit/domain/event_search_index"

	"github.com/google/uuid"
)

// CreateEventSearchIndex 创建事件搜索索引
func (s *Service) CreateEventSearchIndex(ctx context.Context, cmd eventSearchIndexCommands.CreateEventSearchIndexCmd) (uuid.UUID, error) {
	// Parse occurred at
	occurredAt, err := time.Parse(time.RFC3339, cmd.OccurredAt)
	if err != nil {
		return uuid.Nil, err
	}

	// Create domain entity
	eventSearchIndex, err := eventSearchIndexDomain.NewEventSearchIndex(eventSearchIndexDomain.NewEventSearchIndexParams{
		EventID:            cmd.EventID,
		TenantID:           cmd.TenantID,
		AppID:              cmd.AppID,
		ActorType:          cmd.ActorType,
		ActorID:            cmd.ActorID,
		Action:             cmd.Action,
		TargetType:         cmd.TargetType,
		TargetID:           cmd.TargetID,
		Result:             cmd.Result,
		OccurredAt:         occurredAt,
		IP:                 cmd.IP,
		RiskLevel:          cmd.RiskLevel,
		DataClassification: cmd.DataClassification,
		Tags:               cmd.Tags,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.eventSearchIndexRepo.Create.New(ctx, eventSearchIndex); err != nil {
		return uuid.Nil, err
	}

	return eventSearchIndex.ID(), nil
}
