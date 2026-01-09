package mapper

import (
	"nfxid/modules/tenants/domain/member_roles"
	"nfxid/modules/tenants/infrastructure/rdb/models"
)

// MemberRoleDomainToModel 将 Domain MemberRole 转换为 Model MemberRole
func MemberRoleDomainToModel(mr *member_roles.MemberRole) *models.MemberRole {
	if mr == nil {
		return nil
	}

	return &models.MemberRole{
		ID:           mr.ID(),
		TenantID:     mr.TenantID(),
		MemberID:     mr.MemberID(),
		RoleID:       mr.RoleID(),
		AssignedAt:   mr.AssignedAt(),
		AssignedBy:   mr.AssignedBy(),
		ExpiresAt:    mr.ExpiresAt(),
		Scope:        mr.Scope(),
		RevokedAt:    mr.RevokedAt(),
		RevokedBy:    mr.RevokedBy(),
		RevokeReason: mr.RevokeReason(),
	}
}

// MemberRoleModelToDomain 将 Model MemberRole 转换为 Domain MemberRole
func MemberRoleModelToDomain(m *models.MemberRole) *member_roles.MemberRole {
	if m == nil {
		return nil
	}

	state := member_roles.MemberRoleState{
		ID:           m.ID,
		TenantID:     m.TenantID,
		MemberID:     m.MemberID,
		RoleID:       m.RoleID,
		AssignedAt:   m.AssignedAt,
		AssignedBy:   m.AssignedBy,
		ExpiresAt:    m.ExpiresAt,
		Scope:        m.Scope,
		RevokedAt:    m.RevokedAt,
		RevokedBy:    m.RevokedBy,
		RevokeReason: m.RevokeReason,
	}

	return member_roles.NewMemberRoleFromState(state)
}

// MemberRoleModelToUpdates 将 Model MemberRole 转换为更新字段映射
func MemberRoleModelToUpdates(m *models.MemberRole) map[string]any {
	return map[string]any{
		models.MemberRoleCols.TenantID:     m.TenantID,
		models.MemberRoleCols.MemberID:     m.MemberID,
		models.MemberRoleCols.RoleID:       m.RoleID,
		models.MemberRoleCols.AssignedAt:   m.AssignedAt,
		models.MemberRoleCols.AssignedBy:   m.AssignedBy,
		models.MemberRoleCols.ExpiresAt:    m.ExpiresAt,
		models.MemberRoleCols.Scope:        m.Scope,
		models.MemberRoleCols.RevokedAt:    m.RevokedAt,
		models.MemberRoleCols.RevokedBy:    m.RevokedBy,
		models.MemberRoleCols.RevokeReason: m.RevokeReason,
	}
}
