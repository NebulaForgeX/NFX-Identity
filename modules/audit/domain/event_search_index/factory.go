package event_search_index

import (
	"time"

	"github.com/google/uuid"
)

type NewEventSearchIndexParams struct {
	EventID           string
	TenantID          *uuid.UUID
	AppID             *uuid.UUID
	ActorType         ActorType
	ActorID           uuid.UUID
	Action            string
	TargetType        *string
	TargetID          *uuid.UUID
	Result            ResultType
	OccurredAt        time.Time
	IP                *string
	RiskLevel         RiskLevel
	DataClassification DataClassification
	Tags              []string
}

func NewEventSearchIndex(p NewEventSearchIndexParams) (*EventSearchIndex, error) {
	if err := validateEventSearchIndexParams(p); err != nil {
		return nil, err
	}

	riskLevel := p.RiskLevel
	if riskLevel == "" {
		riskLevel = RiskLevelLow
	}

	dataClassification := p.DataClassification
	if dataClassification == "" {
		dataClassification = DataClassificationInternal
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewEventSearchIndexFromState(EventSearchIndexState{
		ID:                 id,
		EventID:            p.EventID,
		TenantID:           p.TenantID,
		AppID:              p.AppID,
		ActorType:          p.ActorType,
		ActorID:            p.ActorID,
		Action:             p.Action,
		TargetType:         p.TargetType,
		TargetID:           p.TargetID,
		Result:             p.Result,
		OccurredAt:         p.OccurredAt,
		IP:                 p.IP,
		RiskLevel:          riskLevel,
		DataClassification: dataClassification,
		Tags:               p.Tags,
		CreatedAt:          now,
	}), nil
}

func NewEventSearchIndexFromState(st EventSearchIndexState) *EventSearchIndex {
	return &EventSearchIndex{state: st}
}

func validateEventSearchIndexParams(p NewEventSearchIndexParams) error {
	if p.EventID == "" {
		return ErrEventIDRequired
	}
	if p.ActorType == "" {
		return ErrActorTypeRequired
	}
	validActorTypes := map[ActorType]struct{}{
		ActorTypeUser:    {},
		ActorTypeService: {},
		ActorTypeSystem:  {},
		ActorTypeAdmin:   {},
	}
	if _, ok := validActorTypes[p.ActorType]; !ok {
		return ErrInvalidActorType
	}
	if p.ActorID == uuid.Nil {
		return ErrActorIDRequired
	}
	if p.Action == "" {
		return ErrActionRequired
	}
	if p.Result == "" {
		return ErrResultRequired
	}
	validResultTypes := map[ResultType]struct{}{
		ResultTypeSuccess: {},
		ResultTypeFailure: {},
		ResultTypeDeny:    {},
		ResultTypeError:   {},
	}
	if _, ok := validResultTypes[p.Result]; !ok {
		return ErrInvalidResultType
	}
	if p.RiskLevel != "" {
		validRiskLevels := map[RiskLevel]struct{}{
			RiskLevelLow:      {},
			RiskLevelMedium:   {},
			RiskLevelHigh:     {},
			RiskLevelCritical: {},
		}
		if _, ok := validRiskLevels[p.RiskLevel]; !ok {
			return ErrInvalidRiskLevel
		}
	}
	if p.DataClassification != "" {
		validClassifications := map[DataClassification]struct{}{
			DataClassificationPublic:       {},
			DataClassificationInternal:     {},
			DataClassificationConfidential: {},
			DataClassificationRestricted:   {},
		}
		if _, ok := validClassifications[p.DataClassification]; !ok {
			return ErrInvalidDataClassification
		}
	}
	return nil
}
