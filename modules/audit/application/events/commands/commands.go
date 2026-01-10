package commands

import (
	"nfxid/modules/audit/domain/events"

	"github.com/google/uuid"
)

// CreateEventCmd 创建事件命令
type CreateEventCmd struct {
	EventID              string
	OccurredAt           string
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
}

// DeleteEventCmd 删除事件命令
type DeleteEventCmd struct {
	EventID uuid.UUID
}
