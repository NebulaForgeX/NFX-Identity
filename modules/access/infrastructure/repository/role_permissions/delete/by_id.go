package delete

import (
	"context"
	"nfxid/modules/access/domain/role_permissions"
	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 RolePermission，实现 role_permissions.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.RolePermission{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return role_permissions.ErrRolePermissionNotFound
	}
	return nil
}
