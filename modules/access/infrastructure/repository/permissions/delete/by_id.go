package delete

import (
	"context"
	"nfxid/modules/access/domain/permissions"
	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 Permission，实现 permissions.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Permission{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return permissions.ErrPermissionNotFound
	}
	return nil
}
