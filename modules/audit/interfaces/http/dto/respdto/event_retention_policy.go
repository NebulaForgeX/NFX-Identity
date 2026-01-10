package respdto

import (
	"time"

	eventRetentionPolicyAppResult "nfxid/modules/audit/application/event_retention_policies/results"

	"github.com/google/uuid"
)

type EventRetentionPolicyDTO struct {
	ID                 uuid.UUID  `json:"id"`
	PolicyName         string     `json:"policy_name"`
	TenantID           *uuid.UUID `json:"tenant_id,omitempty"`
	ActionPattern      *string    `json:"action_pattern,omitempty"`
	DataClassification *string    `json:"data_classification,omitempty"`
	RiskLevel          *string    `json:"risk_level,omitempty"`
	RetentionDays      int        `json:"retention_days"`
	RetentionAction    string     `json:"retention_action"`
	ArchiveLocation    *string    `json:"archive_location,omitempty"`
	Status             string     `json:"status"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	CreatedBy          *uuid.UUID `json:"created_by,omitempty"`
}

// EventRetentionPolicyROToDTO converts application EventRetentionPolicyRO to response DTO
func EventRetentionPolicyROToDTO(v *eventRetentionPolicyAppResult.EventRetentionPolicyRO) *EventRetentionPolicyDTO {
	if v == nil {
		return nil
	}

	dto := &EventRetentionPolicyDTO{
		ID:              v.ID,
		PolicyName:      v.PolicyName,
		TenantID:        v.TenantID,
		ActionPattern:   v.ActionPattern,
		RetentionDays:   v.RetentionDays,
		ArchiveLocation: v.ArchiveLocation,
		Status:          v.Status,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
		CreatedBy:       v.CreatedBy,
	}

	if v.DataClassification != nil {
		dc := string(*v.DataClassification)
		dto.DataClassification = &dc
	}
	if v.RiskLevel != nil {
		rl := string(*v.RiskLevel)
		dto.RiskLevel = &rl
	}
	if v.RetentionAction != "" {
		dto.RetentionAction = string(v.RetentionAction)
	}

	return dto
}
