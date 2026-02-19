package reqdto

import (
	eventSearchIndexAppCommands "nfxid/modules/audit/application/event_search_index/commands"
	eventSearchIndexDomain "nfxid/modules/audit/domain/event_search_index"

	"github.com/google/uuid"
)

type EventSearchIndexCreateRequestDTO struct {
	EventID            string     `json:"event_id" validate:"required"`
	TenantID           *uuid.UUID `json:"tenant_id,omitempty"`
	AppID              *uuid.UUID `json:"app_id,omitempty"`
	ActorType          string     `json:"actor_type" validate:"required"`
	ActorID            uuid.UUID  `json:"actor_id" validate:"required"`
	Action             string     `json:"action" validate:"required"`
	TargetType         *string    `json:"target_type,omitempty"`
	TargetID           *uuid.UUID `json:"target_id,omitempty"`
	Result             string     `json:"result" validate:"required"`
	OccurredAt         string     `json:"occurred_at" validate:"required"`
	IP                 *string    `json:"ip,omitempty"`
	RiskLevel          string     `json:"risk_level" validate:"required"`
	DataClassification string     `json:"data_classification" validate:"required"`
	Tags               []string   `json:"tags,omitempty"`
}

type EventSearchIndexByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

func (r *EventSearchIndexCreateRequestDTO) ToCreateCmd() eventSearchIndexAppCommands.CreateEventSearchIndexCmd {
	cmd := eventSearchIndexAppCommands.CreateEventSearchIndexCmd{
		EventID:    r.EventID,
		TenantID:   r.TenantID,
		AppID:      r.AppID,
		ActorID:    r.ActorID,
		Action:     r.Action,
		TargetType: r.TargetType,
		TargetID:   r.TargetID,
		OccurredAt: r.OccurredAt,
		IP:         r.IP,
		Tags:       r.Tags,
	}

	if r.ActorType != "" {
		cmd.ActorType = eventSearchIndexDomain.ActorType(r.ActorType)
	}
	if r.Result != "" {
		cmd.Result = eventSearchIndexDomain.ResultType(r.Result)
	}
	if r.RiskLevel != "" {
		cmd.RiskLevel = eventSearchIndexDomain.RiskLevel(r.RiskLevel)
	}
	if r.DataClassification != "" {
		cmd.DataClassification = eventSearchIndexDomain.DataClassification(r.DataClassification)
	}

	return cmd
}
