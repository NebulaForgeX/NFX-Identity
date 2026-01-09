package mapper

import (
	"nfxid/modules/access/domain/scope_permissions"
	"nfxid/modules/access/infrastructure/rdb/models"
)

// ScopePermissionDomainToModel 将 Domain ScopePermission 转换为 Model ScopePermission
func ScopePermissionDomainToModel(sp *scope_permissions.ScopePermission) *models.ScopePermission {
	if sp == nil {
		return nil
	}

	return &models.ScopePermission{
		ID:           sp.ID(),
		Scope:        sp.Scope(),
		PermissionID: sp.PermissionID(),
		CreatedAt:    sp.CreatedAt(),
	}
}

// ScopePermissionModelToDomain 将 Model ScopePermission 转换为 Domain ScopePermission
func ScopePermissionModelToDomain(m *models.ScopePermission) *scope_permissions.ScopePermission {
	if m == nil {
		return nil
	}

	state := scope_permissions.ScopePermissionState{
		ID:           m.ID,
		Scope:        m.Scope,
		PermissionID: m.PermissionID,
		CreatedAt:    m.CreatedAt,
	}

	return scope_permissions.NewScopePermissionFromState(state)
}

// ScopePermissionModelToUpdates 将 Model ScopePermission 转换为更新字段映射
func ScopePermissionModelToUpdates(m *models.ScopePermission) map[string]any {
	return map[string]any{
		models.ScopePermissionCols.Scope:        m.Scope,
		models.ScopePermissionCols.PermissionID: m.PermissionID,
	}
}
