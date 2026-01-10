package respdto

import (
	"time"

	eventSearchIndexAppResult "nfxid/modules/audit/application/event_search_index/results"

	"github.com/google/uuid"
)

type EventSearchIndexDTO struct {
	ID                uuid.UUID  `json:"id"`
	EventID           string      `json:"event_id"`
	TenantID          *uuid.UUID  `json:"tenant_id,omitempty"`
	AppID             *uuid.UUID  `json:"app_id,omitempty"`
	ActorType         string      `json:"actor_type"`
	ActorID           uuid.UUID   `json:"actor_id"`
	Action            string      `json:"action"`
	TargetType        *string     `json:"target_type,omitempty"`
	TargetID          *uuid.UUID  `json:"target_id,omitempty"`
	Result            string      `json:"result"`
	OccurredAt        time.Time   `json:"occurred_at"`
	IP                *string     `json:"ip,omitempty"`
	RiskLevel         string      `json:"risk_level"`
	DataClassification string     `json:"data_classification"`
	Tags              []string    `json:"tags,omitempty"`
	CreatedAt         time.Time   `json:"created_at"`
}

// EventSearchIndexROToDTO converts application EventSearchIndexRO to response DTO
func EventSearchIndexROToDTO(v *eventSearchIndexAppResult.EventSearchIndexRO) *EventSearchIndexDTO {
	if v == nil {
		return nil
	}

	return &EventSearchIndexDTO{
		ID:                v.ID,
		EventID:           v.EventID,
		TenantID:          v.TenantID,
		AppID:             v.AppID,
		ActorType:         string(v.ActorType),
		ActorID:           v.ActorID,
		Action:            v.Action,
		TargetType:        v.TargetType,
		TargetID:          v.TargetID,
		Result:            string(v.Result),
		OccurredAt:        v.OccurredAt,
		IP:                v.IP,
		RiskLevel:         string(v.RiskLevel),
		DataClassification: string(v.DataClassification),
		Tags:              v.Tags,
		CreatedAt:         v.CreatedAt,
	}
}
