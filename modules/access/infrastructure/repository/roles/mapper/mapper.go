package mapper

import (
	"nfxid/enums"
	"nfxid/modules/access/domain/roles"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

// RoleDomainToModel 将 Domain Role 转换为 Model Role
func RoleDomainToModel(r *roles.Role) *models.Role {
	if r == nil {
		return nil
	}

	return &models.Role{
		ID:          r.ID(),
		Key:         r.Key(),
		Name:        r.Name(),
		Description: r.Description(),
		ScopeType:   scopeTypeDomainToEnum(r.ScopeType()),
		IsSystem:    r.IsSystem(),
		CreatedAt:   r.CreatedAt(),
		UpdatedAt:   r.UpdatedAt(),
		DeletedAt:   timex.TimeToGormDeletedAt(r.DeletedAt()),
	}
}

// RoleModelToDomain 将 Model Role 转换为 Domain Role
func RoleModelToDomain(m *models.Role) *roles.Role {
	if m == nil {
		return nil
	}

	state := roles.RoleState{
		ID:          m.ID,
		Key:         m.Key,
		Name:        m.Name,
		Description: m.Description,
		ScopeType:   scopeTypeEnumToDomain(m.ScopeType),
		IsSystem:    m.IsSystem,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return roles.NewRoleFromState(state)
}

// RoleModelToUpdates 将 Model Role 转换为更新字段映射
func RoleModelToUpdates(m *models.Role) map[string]any {
	return map[string]any{
		models.RoleCols.Key:         m.Key,
		models.RoleCols.Name:        m.Name,
		models.RoleCols.Description: m.Description,
		models.RoleCols.ScopeType:   m.ScopeType,
		models.RoleCols.IsSystem:    m.IsSystem,
		models.RoleCols.UpdatedAt:   m.UpdatedAt,
		models.RoleCols.DeletedAt:   m.DeletedAt,
	}
}

// 枚举转换辅助函数

func scopeTypeDomainToEnum(st roles.ScopeType) enums.AccessScopeType {
	switch st {
	case roles.ScopeTypeTenant:
		return enums.AccessScopeTypeTenant
	case roles.ScopeTypeApp:
		return enums.AccessScopeTypeApp
	case roles.ScopeTypeGlobal:
		return enums.AccessScopeTypeGlobal
	default:
		return enums.AccessScopeTypeTenant
	}
}

func scopeTypeEnumToDomain(st enums.AccessScopeType) roles.ScopeType {
	switch st {
	case enums.AccessScopeTypeTenant:
		return roles.ScopeTypeTenant
	case enums.AccessScopeTypeApp:
		return roles.ScopeTypeApp
	case enums.AccessScopeTypeGlobal:
		return roles.ScopeTypeGlobal
	default:
		return roles.ScopeTypeTenant
	}
}
