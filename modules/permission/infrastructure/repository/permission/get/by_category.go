package get

import (
	"context"
	permissionDomain "nfxid/modules/permission/domain/permission"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/modules/permission/infrastructure/repository/mapper"
)

// ByCategory 根据 Category 获取 Permission 列表，实现 permissionDomain.Get 接口
func (h *Handler) ByCategory(ctx context.Context, category string) ([]*permissionDomain.Permission, error) {
	var models []models.Permission
	if err := h.db.WithContext(ctx).
		Where("category = ?", category).
		Where("deleted_at IS NULL").
		Order("tag ASC").
		Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*permissionDomain.Permission, len(models))
	for i := range models {
		entities[i] = mapper.PermissionModelToDomain(&models[i])
	}
	return entities, nil
}
