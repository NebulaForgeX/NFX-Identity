package list

import (
	"context"
	"nfxid/modules/permission/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// PermissionTagsByUserID 根据 UserID 获取权限标签列表，实现 userPermissionDomain.List 接口
func (h *Handler) PermissionTagsByUserID(ctx context.Context, userID uuid.UUID) ([]string, error) {
	var tags []string
	err := h.db.WithContext(ctx).
		Model(&models.UserPermission{}).
		Select("p.tag").
		Joins("JOIN permission.permissions p ON p.id = user_permissions.permission_id").
		Where("user_permissions.user_id = ?", userID).
		Where("user_permissions.deleted_at IS NULL").
		Where("p.deleted_at IS NULL").
		Pluck("p.tag", &tags).Error
	return tags, err
}
