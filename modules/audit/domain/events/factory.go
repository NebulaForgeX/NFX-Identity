package events

import (
	"time"

	"github.com/google/uuid"
)

type NewEventParams struct {
	EventID              string
	OccurredAt           time.Time
	TenantID             *uuid.UUID
	AppID                *uuid.UUID
	ActorType            ActorType
	ActorID              uuid.UUID
	ActorTenantMemberID  *uuid.UUID
	Action               string
	TargetType           *string
	TargetID             *uuid.UUID
	Result               ResultType
	FailureReasonCode    *string
	HTTPMethod           *string
	HTTPPath             *string
	HTTPStatus           *int
	RequestID            *string
	TraceID              *string
	IP                   *string
	UserAgent            *string
	GeoCountry           *string
	RiskLevel            RiskLevel
	DataClassification   DataClassification
	PrevHash             *string
	EventHash            *string
	Metadata             map[string]interface{}
}

func NewEvent(p NewEventParams) (*Event, error) {
	if err := validateEventParams(p); err != nil {
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

	now := time.Now().UTC()
	occurredAt := p.OccurredAt
	if occurredAt.IsZero() {
		occurredAt = now
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return NewEventFromState(EventState{
		ID:                   id,
		EventID:              p.EventID,
		OccurredAt:           occurredAt,
		ReceivedAt:           now,
		TenantID:             p.TenantID,
		AppID:                p.AppID,
		ActorType:            p.ActorType,
		ActorID:              p.ActorID,
		ActorTenantMemberID:  p.ActorTenantMemberID,
		Action:               p.Action,
		TargetType:           p.TargetType,
		TargetID:             p.TargetID,
		Result:               p.Result,
		FailureReasonCode:    p.FailureReasonCode,
		HTTPMethod:           p.HTTPMethod,
		HTTPPath:             p.HTTPPath,
		HTTPStatus:           p.HTTPStatus,
		RequestID:            p.RequestID,
		TraceID:              p.TraceID,
		IP:                   p.IP,
		UserAgent:            p.UserAgent,
		GeoCountry:           p.GeoCountry,
		RiskLevel:            riskLevel,
		DataClassification:   dataClassification,
		PrevHash:             p.PrevHash,
		EventHash:            p.EventHash,
		Metadata:             p.Metadata,
		CreatedAt:            now,
	}), nil
}

func NewEventFromState(st EventState) *Event {
	return &Event{state: st}
}

func validateEventParams(p NewEventParams) error {
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
