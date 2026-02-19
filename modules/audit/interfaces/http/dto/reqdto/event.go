package reqdto

import (
	auditAppCommands "nfxid/modules/audit/application/events/commands"
	auditDomain "nfxid/modules/audit/domain/events"

	"github.com/google/uuid"
)

type EventCreateRequestDTO struct {
	EventID             string                 `json:"event_id" validate:"required"`
	OccurredAt          string                 `json:"occurred_at" validate:"required"`
	TenantID            *uuid.UUID             `json:"tenant_id,omitempty"`
	AppID               *uuid.UUID             `json:"app_id,omitempty"`
	ActorType           string                 `json:"actor_type" validate:"required"`
	ActorID             uuid.UUID              `json:"actor_id" validate:"required"`
	ActorTenantMemberID *uuid.UUID             `json:"actor_tenant_member_id,omitempty"`
	Action              string                 `json:"action" validate:"required"`
	TargetType          *string                `json:"target_type,omitempty"`
	TargetID            *uuid.UUID             `json:"target_id,omitempty"`
	Result              string                 `json:"result" validate:"required"`
	FailureReasonCode   *string                `json:"failure_reason_code,omitempty"`
	HTTPMethod          *string                `json:"http_method,omitempty"`
	HTTPPath            *string                `json:"http_path,omitempty"`
	HTTPStatus          *int                   `json:"http_status,omitempty"`
	RequestID           *string                `json:"request_id,omitempty"`
	TraceID             *string                `json:"trace_id,omitempty"`
	IP                  *string                `json:"ip,omitempty"`
	UserAgent           *string                `json:"user_agent,omitempty"`
	GeoCountry          *string                `json:"geo_country,omitempty"`
	RiskLevel           string                 `json:"risk_level,omitempty"`
	DataClassification  string                 `json:"data_classification,omitempty"`
	PrevHash            *string                `json:"prev_hash,omitempty"`
	EventHash           *string                `json:"event_hash,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`
}

type EventByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type EventByEventIDRequestDTO struct {
	EventID string `uri:"event_id" validate:"required"`
}

func (r *EventCreateRequestDTO) ToCreateCmd() auditAppCommands.CreateEventCmd {
	cmd := auditAppCommands.CreateEventCmd{
		EventID:             r.EventID,
		OccurredAt:          r.OccurredAt,
		TenantID:            r.TenantID,
		AppID:               r.AppID,
		ActorID:             r.ActorID,
		ActorTenantMemberID: r.ActorTenantMemberID,
		Action:              r.Action,
		TargetType:          r.TargetType,
		TargetID:            r.TargetID,
		FailureReasonCode:   r.FailureReasonCode,
		HTTPMethod:          r.HTTPMethod,
		HTTPPath:            r.HTTPPath,
		HTTPStatus:          r.HTTPStatus,
		RequestID:           r.RequestID,
		TraceID:             r.TraceID,
		IP:                  r.IP,
		UserAgent:           r.UserAgent,
		GeoCountry:          r.GeoCountry,
		PrevHash:            r.PrevHash,
		EventHash:           r.EventHash,
		Metadata:            r.Metadata,
	}

	if r.ActorType != "" {
		cmd.ActorType = auditDomain.ActorType(r.ActorType)
	}
	if r.Result != "" {
		cmd.Result = auditDomain.ResultType(r.Result)
	}
	if r.RiskLevel != "" {
		cmd.RiskLevel = auditDomain.RiskLevel(r.RiskLevel)
	}
	if r.DataClassification != "" {
		cmd.DataClassification = auditDomain.DataClassification(r.DataClassification)
	}

	return cmd
}
