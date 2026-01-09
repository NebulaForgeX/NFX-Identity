package delete

import (
	"context"
	"nfxid/modules/access/domain/scope_permissions"
	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 ScopePermission，实现 scope_permissions.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.ScopePermission{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return scope_permissions.ErrScopePermissionNotFound
	}
	return nil
}
