package results

import (
	"time"

	"nfxid/modules/audit/domain/event_retention_policies"

	"github.com/google/uuid"
)

type EventRetentionPolicyRO struct {
	ID                 uuid.UUID
	PolicyName         string
	TenantID           *uuid.UUID
	ActionPattern      *string
	DataClassification *event_retention_policies.DataClassification
	RiskLevel          *event_retention_policies.RiskLevel
	RetentionDays      int
	RetentionAction    event_retention_policies.RetentionAction
	ArchiveLocation    *string
	Status             string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	CreatedBy          *uuid.UUID
}

// EventRetentionPolicyMapper 将 Domain EventRetentionPolicy 转换为 Application EventRetentionPolicyRO
func EventRetentionPolicyMapper(erp *event_retention_policies.EventRetentionPolicy) EventRetentionPolicyRO {
	if erp == nil {
		return EventRetentionPolicyRO{}
	}

	return EventRetentionPolicyRO{
		ID:                 erp.ID(),
		PolicyName:         erp.PolicyName(),
		TenantID:           erp.TenantID(),
		ActionPattern:      erp.ActionPattern(),
		DataClassification: erp.DataClassification(),
		RiskLevel:          erp.RiskLevel(),
		RetentionDays:      erp.RetentionDays(),
		RetentionAction:    erp.RetentionAction(),
		ArchiveLocation:    erp.ArchiveLocation(),
		Status:             erp.Status(),
		CreatedAt:          erp.CreatedAt(),
		UpdatedAt:          erp.UpdatedAt(),
		CreatedBy:          erp.CreatedBy(),
	}
}
