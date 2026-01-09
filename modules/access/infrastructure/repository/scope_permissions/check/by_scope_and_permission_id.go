package check

import (
	"context"
	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByScopeAndPermissionID 根据 Scope 和 PermissionID 检查 ScopePermission 是否存在，实现 scope_permissions.Check 接口
func (h *Handler) ByScopeAndPermissionID(ctx context.Context, scope string, permissionID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.ScopePermission{}).
		Where("scope = ? AND permission_id = ?", scope, permissionID).
		Count(&count).Error
	return count > 0, err
}
