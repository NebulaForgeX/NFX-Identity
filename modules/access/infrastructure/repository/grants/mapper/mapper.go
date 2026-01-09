package mapper

import (
	"nfxid/enums"
	"nfxid/modules/access/domain/grants"
	"nfxid/modules/access/infrastructure/rdb/models"
)

// GrantDomainToModel 将 Domain Grant 转换为 Model Grant
func GrantDomainToModel(g *grants.Grant) *models.Grant {
	if g == nil {
		return nil
	}

	return &models.Grant{
		ID:           g.ID(),
		SubjectType:  subjectTypeDomainToEnum(g.SubjectType()),
		SubjectID:    g.SubjectID(),
		GrantType:    grantTypeDomainToEnum(g.GrantType()),
		GrantRefID:   g.GrantRefID(),
		TenantID:     g.TenantID(),
		AppID:        g.AppID(),
		ResourceType: g.ResourceType(),
		ResourceID:   g.ResourceID(),
		Effect:       grantEffectDomainToEnum(g.Effect()),
		ExpiresAt:    g.ExpiresAt(),
		CreatedAt:    g.CreatedAt(),
		CreatedBy:    g.CreatedBy(),
		RevokedAt:    g.RevokedAt(),
		RevokedBy:    g.RevokedBy(),
		RevokeReason: g.RevokeReason(),
	}
}

// GrantModelToDomain 将 Model Grant 转换为 Domain Grant
func GrantModelToDomain(m *models.Grant) *grants.Grant {
	if m == nil {
		return nil
	}

	state := grants.GrantState{
		ID:           m.ID,
		SubjectType:  subjectTypeEnumToDomain(m.SubjectType),
		SubjectID:    m.SubjectID,
		GrantType:    grantTypeEnumToDomain(m.GrantType),
		GrantRefID:   m.GrantRefID,
		TenantID:     m.TenantID,
		AppID:        m.AppID,
		ResourceType: m.ResourceType,
		ResourceID:   m.ResourceID,
		Effect:       grantEffectEnumToDomain(m.Effect),
		ExpiresAt:    m.ExpiresAt,
		CreatedAt:    m.CreatedAt,
		CreatedBy:    m.CreatedBy,
		RevokedAt:    m.RevokedAt,
		RevokedBy:    m.RevokedBy,
		RevokeReason: m.RevokeReason,
	}

	return grants.NewGrantFromState(state)
}

// GrantModelToUpdates 将 Model Grant 转换为更新字段映射
func GrantModelToUpdates(m *models.Grant) map[string]any {
	return map[string]any{
		models.GrantCols.SubjectType:  m.SubjectType,
		models.GrantCols.SubjectID:     m.SubjectID,
		models.GrantCols.GrantType:     m.GrantType,
		models.GrantCols.GrantRefID:    m.GrantRefID,
		models.GrantCols.TenantID:      m.TenantID,
		models.GrantCols.AppID:         m.AppID,
		models.GrantCols.ResourceType: m.ResourceType,
		models.GrantCols.ResourceID:    m.ResourceID,
		models.GrantCols.Effect:        m.Effect,
		models.GrantCols.ExpiresAt:     m.ExpiresAt,
		models.GrantCols.CreatedBy:    m.CreatedBy,
		models.GrantCols.RevokedAt:     m.RevokedAt,
		models.GrantCols.RevokedBy:     m.RevokedBy,
		models.GrantCols.RevokeReason:  m.RevokeReason,
	}
}

// 枚举转换辅助函数

func subjectTypeDomainToEnum(st grants.SubjectType) enums.AccessSubjectType {
	switch st {
	case grants.SubjectTypeUser:
		return enums.AccessSubjectTypeUser
	case grants.SubjectTypeClient:
		return enums.AccessSubjectTypeClient
	default:
		return enums.AccessSubjectTypeUser
	}
}

func subjectTypeEnumToDomain(st enums.AccessSubjectType) grants.SubjectType {
	switch st {
	case enums.AccessSubjectTypeUser:
		return grants.SubjectTypeUser
	case enums.AccessSubjectTypeClient:
		return grants.SubjectTypeClient
	default:
		return grants.SubjectTypeUser
	}
}

func grantTypeDomainToEnum(gt grants.GrantType) enums.AccessGrantType {
	switch gt {
	case grants.GrantTypeRole:
		return enums.AccessGrantTypeRole
	case grants.GrantTypePermission:
		return enums.AccessGrantTypePermission
	default:
		return enums.AccessGrantTypeRole
	}
}

func grantTypeEnumToDomain(gt enums.AccessGrantType) grants.GrantType {
	switch gt {
	case enums.AccessGrantTypeRole:
		return grants.GrantTypeRole
	case enums.AccessGrantTypePermission:
		return grants.GrantTypePermission
	default:
		return grants.GrantTypeRole
	}
}

func grantEffectDomainToEnum(ge grants.GrantEffect) enums.AccessGrantEffect {
	switch ge {
	case grants.GrantEffectAllow:
		return enums.AccessGrantEffectAllow
	case grants.GrantEffectDeny:
		return enums.AccessGrantEffectDeny
	default:
		return enums.AccessGrantEffectAllow
	}
}

func grantEffectEnumToDomain(ge enums.AccessGrantEffect) grants.GrantEffect {
	switch ge {
	case enums.AccessGrantEffectAllow:
		return grants.GrantEffectAllow
	case enums.AccessGrantEffectDeny:
		return grants.GrantEffectDeny
	default:
		return grants.GrantEffectAllow
	}
}
