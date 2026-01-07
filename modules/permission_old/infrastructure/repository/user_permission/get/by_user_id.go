package get

import (
	"context"
	userPermissionDomain "nfxid/modules/permission/domain/user_permission"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/modules/permission/infrastructure/repository/mapper"

	"github.com/google/uuid"
)

// ByUserID 根据 UserID 获取 UserPermission 列表，实现 userPermissionDomain.Get 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) ([]*userPermissionDomain.UserPermission, error) {
	var models []models.UserPermission
	if err := h.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*userPermissionDomain.UserPermission, len(models))
	for i := range models {
		entities[i] = mapper.UserPermissionModelToDomain(&models[i])
	}
	return entities, nil
}
