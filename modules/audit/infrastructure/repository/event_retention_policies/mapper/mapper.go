package mapper

import (
	"nfxid/enums"
	"nfxid/modules/audit/domain/event_retention_policies"
	"nfxid/modules/audit/infrastructure/rdb/models"
)

// EventRetentionPolicyDomainToModel 将 Domain EventRetentionPolicy 转换为 Model EventRetentionPolicy
func EventRetentionPolicyDomainToModel(erp *event_retention_policies.EventRetentionPolicy) *models.EventRetentionPolicy {
	if erp == nil {
		return nil
	}

	var dataClassification *enums.AuditDataClassification
	if erp.DataClassification() != nil {
		dc := enums.AuditDataClassification(*erp.DataClassification())
		dataClassification = &dc
	}

	var riskLevel *enums.AuditRiskLevel
	if erp.RiskLevel() != nil {
		rl := enums.AuditRiskLevel(*erp.RiskLevel())
		riskLevel = &rl
	}

	return &models.EventRetentionPolicy{
		ID:                 erp.ID(),
		PolicyName:         erp.PolicyName(),
		TenantID:           erp.TenantID(),
		ActionPattern:      erp.ActionPattern(),
		DataClassification: dataClassification,
		RiskLevel:          riskLevel,
		RetentionDays:      erp.RetentionDays(),
		RetentionAction:    enums.AuditRetentionAction(erp.RetentionAction()),
		ArchiveLocation:    erp.ArchiveLocation(),
		Status:             erp.Status(),
		CreatedAt:          erp.CreatedAt(),
		UpdatedAt:          erp.UpdatedAt(),
		CreatedBy:          erp.CreatedBy(),
	}
}

// EventRetentionPolicyModelToDomain 将 Model EventRetentionPolicy 转换为 Domain EventRetentionPolicy
func EventRetentionPolicyModelToDomain(m *models.EventRetentionPolicy) *event_retention_policies.EventRetentionPolicy {
	if m == nil {
		return nil
	}

	var dataClassification *event_retention_policies.DataClassification
	if m.DataClassification != nil {
		dc := event_retention_policies.DataClassification(*m.DataClassification)
		dataClassification = &dc
	}

	var riskLevel *event_retention_policies.RiskLevel
	if m.RiskLevel != nil {
		rl := event_retention_policies.RiskLevel(*m.RiskLevel)
		riskLevel = &rl
	}

	state := event_retention_policies.EventRetentionPolicyState{
		ID:                 m.ID,
		PolicyName:         m.PolicyName,
		TenantID:           m.TenantID,
		ActionPattern:      m.ActionPattern,
		DataClassification: dataClassification,
		RiskLevel:          riskLevel,
		RetentionDays:      m.RetentionDays,
		RetentionAction:    event_retention_policies.RetentionAction(m.RetentionAction),
		ArchiveLocation:    m.ArchiveLocation,
		Status:             m.Status,
		CreatedAt:          m.CreatedAt,
		UpdatedAt:          m.UpdatedAt,
		CreatedBy:          m.CreatedBy,
	}

	return event_retention_policies.NewEventRetentionPolicyFromState(state)
}

// EventRetentionPolicyModelToUpdates 将 Model EventRetentionPolicy 转换为更新字段映射
func EventRetentionPolicyModelToUpdates(m *models.EventRetentionPolicy) map[string]any {
	return map[string]any{
		models.EventRetentionPolicyCols.PolicyName:         m.PolicyName,
		models.EventRetentionPolicyCols.TenantID:           m.TenantID,
		models.EventRetentionPolicyCols.ActionPattern:      m.ActionPattern,
		models.EventRetentionPolicyCols.DataClassification: m.DataClassification,
		models.EventRetentionPolicyCols.RiskLevel:          m.RiskLevel,
		models.EventRetentionPolicyCols.RetentionDays:      m.RetentionDays,
		models.EventRetentionPolicyCols.RetentionAction:    m.RetentionAction,
		models.EventRetentionPolicyCols.ArchiveLocation:    m.ArchiveLocation,
		models.EventRetentionPolicyCols.Status:             m.Status,
		models.EventRetentionPolicyCols.UpdatedAt:          m.UpdatedAt,
		models.EventRetentionPolicyCols.CreatedBy:          m.CreatedBy,
	}
}
