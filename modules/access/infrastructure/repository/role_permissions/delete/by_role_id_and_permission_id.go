package delete

import (
	"context"
	"nfxid/modules/access/domain/role_permissions"
	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByRoleIDAndPermissionID 根据 RoleID 和 PermissionID 删除 RolePermission，实现 role_permissions.Delete 接口
func (h *Handler) ByRoleIDAndPermissionID(ctx context.Context, roleID, permissionID uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("role_id = ? AND permission_id = ?", roleID, permissionID).
		Delete(&models.RolePermission{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return role_permissions.ErrRolePermissionNotFound
	}
	return nil
}
