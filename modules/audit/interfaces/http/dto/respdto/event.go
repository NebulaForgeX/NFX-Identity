package respdto

import (
	"time"

	auditAppResult "nfxid/modules/audit/application/events/results"

	"github.com/google/uuid"
)

type EventDTO struct {
	ID                   uuid.UUID              `json:"id"`
	EventID              string                 `json:"event_id"`
	OccurredAt           time.Time              `json:"occurred_at"`
	ReceivedAt           time.Time              `json:"received_at"`
	TenantID             *uuid.UUID             `json:"tenant_id,omitempty"`
	AppID                *uuid.UUID             `json:"app_id,omitempty"`
	ActorType            string                 `json:"actor_type"`
	ActorID              uuid.UUID              `json:"actor_id"`
	ActorTenantMemberID  *uuid.UUID             `json:"actor_tenant_member_id,omitempty"`
	Action               string                 `json:"action"`
	TargetType           *string                `json:"target_type,omitempty"`
	TargetID             *uuid.UUID             `json:"target_id,omitempty"`
	Result               string                 `json:"result"`
	FailureReasonCode    *string                `json:"failure_reason_code,omitempty"`
	HTTPMethod           *string                `json:"http_method,omitempty"`
	HTTPPath             *string                `json:"http_path,omitempty"`
	HTTPStatus           *int                   `json:"http_status,omitempty"`
	RequestID            *string                `json:"request_id,omitempty"`
	TraceID              *string                `json:"trace_id,omitempty"`
	IP                   *string                `json:"ip,omitempty"`
	UserAgent            *string                `json:"user_agent,omitempty"`
	GeoCountry           *string                `json:"geo_country,omitempty"`
	RiskLevel            string                 `json:"risk_level"`
	DataClassification   string                 `json:"data_classification"`
	PrevHash             *string                `json:"prev_hash,omitempty"`
	EventHash            *string                `json:"event_hash,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt            time.Time              `json:"created_at"`
}

// EventROToDTO converts application EventRO to response DTO
func EventROToDTO(v *auditAppResult.EventRO) *EventDTO {
	if v == nil {
		return nil
	}

	return &EventDTO{
		ID:                   v.ID,
		EventID:              v.EventID,
		OccurredAt:           v.OccurredAt,
		ReceivedAt:           v.ReceivedAt,
		TenantID:             v.TenantID,
		AppID:                v.AppID,
		ActorType:            string(v.ActorType),
		ActorID:              v.ActorID,
		ActorTenantMemberID:  v.ActorTenantMemberID,
		Action:               v.Action,
		TargetType:           v.TargetType,
		TargetID:             v.TargetID,
		Result:               string(v.Result),
		FailureReasonCode:    v.FailureReasonCode,
		HTTPMethod:           v.HTTPMethod,
		HTTPPath:             v.HTTPPath,
		HTTPStatus:           v.HTTPStatus,
		RequestID:            v.RequestID,
		TraceID:              v.TraceID,
		IP:                   v.IP,
		UserAgent:            v.UserAgent,
		GeoCountry:           v.GeoCountry,
		RiskLevel:            string(v.RiskLevel),
		DataClassification:   string(v.DataClassification),
		PrevHash:             v.PrevHash,
		EventHash:            v.EventHash,
		Metadata:             v.Metadata,
		CreatedAt:            v.CreatedAt,
	}
}
