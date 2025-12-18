package get

import (
	"context"
	permissionDomain "nfxid/modules/permission/domain/permission"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/modules/permission/infrastructure/repository/mapper"
)

// ByTags 根据 Tags 获取 Permission 列表，实现 permissionDomain.Get 接口
func (h *Handler) ByTags(ctx context.Context, tags []string) ([]*permissionDomain.Permission, error) {
	var models []models.Permission
	if err := h.db.WithContext(ctx).
		Where("tag IN ?", tags).
		Where("deleted_at IS NULL").
		Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*permissionDomain.Permission, len(models))
	for i := range models {
		entities[i] = mapper.PermissionModelToDomain(&models[i])
	}
	return entities, nil
}
