package reqdto

import (
	eventRetentionPolicyAppCommands "nfxid/modules/audit/application/event_retention_policies/commands"
	eventRetentionPolicyDomain "nfxid/modules/audit/domain/event_retention_policies"

	"github.com/google/uuid"
)

type EventRetentionPolicyCreateRequestDTO struct {
	PolicyName        string   `json:"policy_name" validate:"required"`
	TenantID          *uuid.UUID `json:"tenant_id,omitempty"`
	ActionPattern     *string   `json:"action_pattern,omitempty"`
	DataClassification *string   `json:"data_classification,omitempty"`
	RiskLevel         *string   `json:"risk_level,omitempty"`
	RetentionDays     int       `json:"retention_days" validate:"required"`
	RetentionAction   string    `json:"retention_action" validate:"required"`
	ArchiveLocation   *string   `json:"archive_location,omitempty"`
	Status            string    `json:"status,omitempty"`
	CreatedBy         *uuid.UUID `json:"created_by,omitempty"`
}

type EventRetentionPolicyUpdateRequestDTO struct {
	ID                uuid.UUID `params:"id" validate:"required,uuid"`
	ActionPattern     *string   `json:"action_pattern,omitempty"`
	DataClassification *string   `json:"data_classification,omitempty"`
	RiskLevel         *string   `json:"risk_level,omitempty"`
	RetentionDays     int       `json:"retention_days" validate:"required"`
	RetentionAction   string    `json:"retention_action" validate:"required"`
	ArchiveLocation   *string   `json:"archive_location,omitempty"`
}

type EventRetentionPolicyByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

func (r *EventRetentionPolicyCreateRequestDTO) ToCreateCmd() eventRetentionPolicyAppCommands.CreateEventRetentionPolicyCmd {
	cmd := eventRetentionPolicyAppCommands.CreateEventRetentionPolicyCmd{
		PolicyName:      r.PolicyName,
		TenantID:        r.TenantID,
		ActionPattern:   r.ActionPattern,
		RetentionDays:   r.RetentionDays,
		ArchiveLocation: r.ArchiveLocation,
		Status:          r.Status,
		CreatedBy:       r.CreatedBy,
	}

	if r.DataClassification != nil {
		dc := eventRetentionPolicyDomain.DataClassification(*r.DataClassification)
		cmd.DataClassification = &dc
	}
	if r.RiskLevel != nil {
		rl := eventRetentionPolicyDomain.RiskLevel(*r.RiskLevel)
		cmd.RiskLevel = &rl
	}
	if r.RetentionAction != "" {
		cmd.RetentionAction = eventRetentionPolicyDomain.RetentionAction(r.RetentionAction)
	}

	return cmd
}

func (r *EventRetentionPolicyUpdateRequestDTO) ToUpdateCmd() eventRetentionPolicyAppCommands.UpdateEventRetentionPolicyCmd {
	cmd := eventRetentionPolicyAppCommands.UpdateEventRetentionPolicyCmd{
		EventRetentionPolicyID: r.ID,
		ActionPattern:           r.ActionPattern,
		RetentionDays:           r.RetentionDays,
		ArchiveLocation:         r.ArchiveLocation,
	}

	if r.DataClassification != nil {
		dc := eventRetentionPolicyDomain.DataClassification(*r.DataClassification)
		cmd.DataClassification = &dc
	}
	if r.RiskLevel != nil {
		rl := eventRetentionPolicyDomain.RiskLevel(*r.RiskLevel)
		cmd.RiskLevel = &rl
	}
	if r.RetentionAction != "" {
		cmd.RetentionAction = eventRetentionPolicyDomain.RetentionAction(r.RetentionAction)
	}

	return cmd
}
