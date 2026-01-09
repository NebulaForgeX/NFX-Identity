package check

import (
	"context"
	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByRoleIDAndPermissionID 根据 RoleID 和 PermissionID 检查 RolePermission 是否存在，实现 role_permissions.Check 接口
func (h *Handler) ByRoleIDAndPermissionID(ctx context.Context, roleID, permissionID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.RolePermission{}).
		Where("role_id = ? AND permission_id = ?", roleID, permissionID).
		Count(&count).Error
	return count > 0, err
}
