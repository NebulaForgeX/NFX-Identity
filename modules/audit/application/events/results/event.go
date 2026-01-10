package results

import (
	"time"

	"nfxid/modules/audit/domain/events"

	"github.com/google/uuid"
)

type EventRO struct {
	ID                   uuid.UUID
	EventID              string
	OccurredAt           time.Time
	ReceivedAt           time.Time
	TenantID             *uuid.UUID
	AppID                *uuid.UUID
	ActorType            events.ActorType
	ActorID              uuid.UUID
	ActorTenantMemberID  *uuid.UUID
	Action               string
	TargetType           *string
	TargetID             *uuid.UUID
	Result               events.ResultType
	FailureReasonCode    *string
	HTTPMethod           *string
	HTTPPath             *string
	HTTPStatus           *int
	RequestID            *string
	TraceID              *string
	IP                   *string
	UserAgent            *string
	GeoCountry           *string
	RiskLevel            events.RiskLevel
	DataClassification   events.DataClassification
	PrevHash             *string
	EventHash            *string
	Metadata             map[string]interface{}
	CreatedAt            time.Time
}

// EventMapper 将 Domain Event 转换为 Application EventRO
func EventMapper(e *events.Event) EventRO {
	if e == nil {
		return EventRO{}
	}

	return EventRO{
		ID:                   e.ID(),
		EventID:              e.EventID(),
		OccurredAt:           e.OccurredAt(),
		ReceivedAt:           e.ReceivedAt(),
		TenantID:             e.TenantID(),
		AppID:                e.AppID(),
		ActorType:            e.ActorType(),
		ActorID:              e.ActorID(),
		ActorTenantMemberID:  e.ActorTenantMemberID(),
		Action:               e.Action(),
		TargetType:           e.TargetType(),
		TargetID:             e.TargetID(),
		Result:               e.Result(),
		FailureReasonCode:    e.FailureReasonCode(),
		HTTPMethod:           e.HTTPMethod(),
		HTTPPath:             e.HTTPPath(),
		HTTPStatus:           e.HTTPStatus(),
		RequestID:            e.RequestID(),
		TraceID:              e.TraceID(),
		IP:                   e.IP(),
		UserAgent:            e.UserAgent(),
		GeoCountry:           e.GeoCountry(),
		RiskLevel:            e.RiskLevel(),
		DataClassification:   e.DataClassification(),
		PrevHash:             e.PrevHash(),
		EventHash:            e.EventHash(),
		Metadata:             e.Metadata(),
		CreatedAt:            e.CreatedAt(),
	}
}
