package delete

import (
	"context"
	permissionDomainErrors "nfxid/modules/permission/domain/permission/errors"
	"nfxid/modules/permission/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 Permission，实现 permissionDomain.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Permission{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return permissionDomainErrors.ErrPermissionNotFound
	}
	return nil
}
