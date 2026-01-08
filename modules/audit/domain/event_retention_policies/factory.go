package event_retention_policies

import (
	"time"

	"github.com/google/uuid"
)

type NewEventRetentionPolicyParams struct {
	PolicyName        string
	TenantID          *uuid.UUID
	ActionPattern     *string
	DataClassification *DataClassification
	RiskLevel         *RiskLevel
	RetentionDays     int
	RetentionAction   RetentionAction
	ArchiveLocation   *string
	Status            string
	CreatedBy         *uuid.UUID
}

func NewEventRetentionPolicy(p NewEventRetentionPolicyParams) (*EventRetentionPolicy, error) {
	if err := validateEventRetentionPolicyParams(p); err != nil {
		return nil, err
	}

	retentionAction := p.RetentionAction
	if retentionAction == "" {
		retentionAction = RetentionActionArchive
	}

	status := p.Status
	if status == "" {
		status = "active"
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewEventRetentionPolicyFromState(EventRetentionPolicyState{
		ID:                 id,
		PolicyName:         p.PolicyName,
		TenantID:           p.TenantID,
		ActionPattern:      p.ActionPattern,
		DataClassification: p.DataClassification,
		RiskLevel:          p.RiskLevel,
		RetentionDays:      p.RetentionDays,
		RetentionAction:    retentionAction,
		ArchiveLocation:    p.ArchiveLocation,
		Status:             status,
		CreatedAt:          now,
		UpdatedAt:          now,
		CreatedBy:          p.CreatedBy,
	}), nil
}

func NewEventRetentionPolicyFromState(st EventRetentionPolicyState) *EventRetentionPolicy {
	return &EventRetentionPolicy{state: st}
}

func validateEventRetentionPolicyParams(p NewEventRetentionPolicyParams) error {
	if p.PolicyName == "" {
		return ErrPolicyNameRequired
	}
	if p.RetentionDays <= 0 {
		return ErrRetentionDaysRequired
	}
	if p.RetentionAction != "" {
		validActions := map[RetentionAction]struct{}{
			RetentionActionArchive: {},
			RetentionActionDelete:  {},
			RetentionActionExport:  {},
		}
		if _, ok := validActions[p.RetentionAction]; !ok {
			return ErrInvalidRetentionAction
		}
	}
	return nil
}
