package commands

import (
	"nfxid/modules/audit/domain/event_retention_policies"

	"github.com/google/uuid"
)

// CreateEventRetentionPolicyCmd 创建事件保留策略命令
type CreateEventRetentionPolicyCmd struct {
	PolicyName        string
	TenantID          *uuid.UUID
	ActionPattern     *string
	DataClassification *event_retention_policies.DataClassification
	RiskLevel         *event_retention_policies.RiskLevel
	RetentionDays     int
	RetentionAction   event_retention_policies.RetentionAction
	ArchiveLocation   *string
	Status            string
	CreatedBy         *uuid.UUID
}

// UpdateEventRetentionPolicyCmd 更新事件保留策略命令
type UpdateEventRetentionPolicyCmd struct {
	EventRetentionPolicyID uuid.UUID
	ActionPattern          *string
	DataClassification     *event_retention_policies.DataClassification
	RiskLevel              *event_retention_policies.RiskLevel
	RetentionDays          int
	RetentionAction        event_retention_policies.RetentionAction
	ArchiveLocation        *string
}

// UpdateEventRetentionPolicyStatusCmd 更新事件保留策略状态命令
type UpdateEventRetentionPolicyStatusCmd struct {
	EventRetentionPolicyID uuid.UUID
	Status                 string
}

// DeleteEventRetentionPolicyCmd 删除事件保留策略命令
type DeleteEventRetentionPolicyCmd struct {
	EventRetentionPolicyID uuid.UUID
}
