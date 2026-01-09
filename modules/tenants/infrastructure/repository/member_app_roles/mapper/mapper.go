package mapper

import (
	"nfxid/modules/tenants/domain/member_app_roles"
	"nfxid/modules/tenants/infrastructure/rdb/models"
)

// MemberAppRoleDomainToModel 将 Domain MemberAppRole 转换为 Model MemberAppRole
func MemberAppRoleDomainToModel(mar *member_app_roles.MemberAppRole) *models.MemberAppRole {
	if mar == nil {
		return nil
	}

	return &models.MemberAppRole{
		ID:           mar.ID(),
		MemberID:     mar.MemberID(),
		AppID:        mar.AppID(),
		RoleID:       mar.RoleID(),
		AssignedAt:   mar.AssignedAt(),
		AssignedBy:   mar.AssignedBy(),
		ExpiresAt:    mar.ExpiresAt(),
		RevokedAt:    mar.RevokedAt(),
		RevokedBy:    mar.RevokedBy(),
		RevokeReason: mar.RevokeReason(),
	}
}

// MemberAppRoleModelToDomain 将 Model MemberAppRole 转换为 Domain MemberAppRole
func MemberAppRoleModelToDomain(m *models.MemberAppRole) *member_app_roles.MemberAppRole {
	if m == nil {
		return nil
	}

	state := member_app_roles.MemberAppRoleState{
		ID:           m.ID,
		MemberID:     m.MemberID,
		AppID:        m.AppID,
		RoleID:       m.RoleID,
		AssignedAt:   m.AssignedAt,
		AssignedBy:   m.AssignedBy,
		ExpiresAt:    m.ExpiresAt,
		RevokedAt:    m.RevokedAt,
		RevokedBy:    m.RevokedBy,
		RevokeReason: m.RevokeReason,
	}

	return member_app_roles.NewMemberAppRoleFromState(state)
}

// MemberAppRoleModelToUpdates 将 Model MemberAppRole 转换为更新字段映射
func MemberAppRoleModelToUpdates(m *models.MemberAppRole) map[string]any {
	return map[string]any{
		models.MemberAppRoleCols.MemberID:     m.MemberID,
		models.MemberAppRoleCols.AppID:        m.AppID,
		models.MemberAppRoleCols.RoleID:       m.RoleID,
		models.MemberAppRoleCols.AssignedAt:   m.AssignedAt,
		models.MemberAppRoleCols.AssignedBy:   m.AssignedBy,
		models.MemberAppRoleCols.ExpiresAt:    m.ExpiresAt,
		models.MemberAppRoleCols.RevokedAt:    m.RevokedAt,
		models.MemberAppRoleCols.RevokedBy:    m.RevokedBy,
		models.MemberAppRoleCols.RevokeReason: m.RevokeReason,
	}
}
