package delete

import (
	"context"
	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByRoleID 根据 RoleID 删除所有 RolePermission，实现 role_permissions.Delete 接口
func (h *Handler) ByRoleID(ctx context.Context, roleID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where("role_id = ?", roleID).
		Delete(&models.RolePermission{}).Error
}
