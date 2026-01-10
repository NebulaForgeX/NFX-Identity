package results

import (
	"time"

	"nfxid/modules/audit/domain/event_search_index"

	"github.com/google/uuid"
)

type EventSearchIndexRO struct {
	ID                uuid.UUID
	EventID           string
	TenantID          *uuid.UUID
	AppID             *uuid.UUID
	ActorType         event_search_index.ActorType
	ActorID           uuid.UUID
	Action            string
	TargetType        *string
	TargetID          *uuid.UUID
	Result            event_search_index.ResultType
	OccurredAt        time.Time
	IP                *string
	RiskLevel         event_search_index.RiskLevel
	DataClassification event_search_index.DataClassification
	Tags              []string
	CreatedAt         time.Time
}

// EventSearchIndexMapper 将 Domain EventSearchIndex 转换为 Application EventSearchIndexRO
func EventSearchIndexMapper(esi *event_search_index.EventSearchIndex) EventSearchIndexRO {
	if esi == nil {
		return EventSearchIndexRO{}
	}

	return EventSearchIndexRO{
		ID:                esi.ID(),
		EventID:           esi.EventID(),
		TenantID:          esi.TenantID(),
		AppID:             esi.AppID(),
		ActorType:         esi.ActorType(),
		ActorID:           esi.ActorID(),
		Action:            esi.Action(),
		TargetType:        esi.TargetType(),
		TargetID:          esi.TargetID(),
		Result:            esi.Result(),
		OccurredAt:        esi.OccurredAt(),
		IP:                esi.IP(),
		RiskLevel:         esi.RiskLevel(),
		DataClassification: esi.DataClassification(),
		Tags:              esi.Tags(),
		CreatedAt:         esi.CreatedAt(),
	}
}
