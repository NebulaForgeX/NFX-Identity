package mapper

import (
	"nfxid/modules/access/domain/role_permissions"
	"nfxid/modules/access/infrastructure/rdb/models"
)

// RolePermissionDomainToModel 将 Domain RolePermission 转换为 Model RolePermission
func RolePermissionDomainToModel(rp *role_permissions.RolePermission) *models.RolePermission {
	if rp == nil {
		return nil
	}

	return &models.RolePermission{
		ID:           rp.ID(),
		RoleID:       rp.RoleID(),
		PermissionID: rp.PermissionID(),
		CreatedAt:    rp.CreatedAt(),
		CreatedBy:    rp.CreatedBy(),
	}
}

// RolePermissionModelToDomain 将 Model RolePermission 转换为 Domain RolePermission
func RolePermissionModelToDomain(m *models.RolePermission) *role_permissions.RolePermission {
	if m == nil {
		return nil
	}

	state := role_permissions.RolePermissionState{
		ID:           m.ID,
		RoleID:       m.RoleID,
		PermissionID: m.PermissionID,
		CreatedAt:    m.CreatedAt,
		CreatedBy:    m.CreatedBy,
	}

	return role_permissions.NewRolePermissionFromState(state)
}

// RolePermissionModelToUpdates 将 Model RolePermission 转换为更新字段映射
func RolePermissionModelToUpdates(m *models.RolePermission) map[string]any {
	return map[string]any{
		models.RolePermissionCols.RoleID:       m.RoleID,
		models.RolePermissionCols.PermissionID: m.PermissionID,
		models.RolePermissionCols.CreatedBy:    m.CreatedBy,
	}
}
