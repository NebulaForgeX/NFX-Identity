package delete

import (
	"context"
	"nfxid/modules/access/domain/scope_permissions"
	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByScopeAndPermissionID 根据 Scope 和 PermissionID 删除 ScopePermission，实现 scope_permissions.Delete 接口
func (h *Handler) ByScopeAndPermissionID(ctx context.Context, scope string, permissionID uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("scope = ? AND permission_id = ?", scope, permissionID).
		Delete(&models.ScopePermission{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return scope_permissions.ErrScopePermissionNotFound
	}
	return nil
}
