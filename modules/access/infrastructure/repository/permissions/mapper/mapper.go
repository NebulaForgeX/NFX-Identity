package mapper

import (
	"nfxid/modules/access/domain/permissions"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"
)

// PermissionDomainToModel 将 Domain Permission 转换为 Model Permission
func PermissionDomainToModel(p *permissions.Permission) *models.Permission {
	if p == nil {
		return nil
	}

	return &models.Permission{
		ID:          p.ID(),
		Key:         p.Key(),
		Name:        p.Name(),
		Description: p.Description(),
		IsSystem:    p.IsSystem(),
		CreatedAt:   p.CreatedAt(),
		UpdatedAt:   p.UpdatedAt(),
		DeletedAt:   timex.TimeToGormDeletedAt(p.DeletedAt()),
	}
}

// PermissionModelToDomain 将 Model Permission 转换为 Domain Permission
func PermissionModelToDomain(m *models.Permission) *permissions.Permission {
	if m == nil {
		return nil
	}

	state := permissions.PermissionState{
		ID:          m.ID,
		Key:         m.Key,
		Name:        m.Name,
		Description: m.Description,
		IsSystem:    m.IsSystem,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return permissions.NewPermissionFromState(state)
}

// PermissionModelToUpdates 将 Model Permission 转换为更新字段映射
func PermissionModelToUpdates(m *models.Permission) map[string]any {
	return map[string]any{
		models.PermissionCols.Key:         m.Key,
		models.PermissionCols.Name:        m.Name,
		models.PermissionCols.Description: m.Description,
		models.PermissionCols.IsSystem:    m.IsSystem,
		models.PermissionCols.UpdatedAt:   m.UpdatedAt,
		models.PermissionCols.DeletedAt:   m.DeletedAt,
	}
}
