package delete

import (
	"context"
	userPermissionDomainErrors "nfxid/modules/permission/domain/user_permission/errors"
	"nfxid/modules/permission/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 UserPermission，实现 userPermissionDomain.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.UserPermission{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return userPermissionDomainErrors.ErrUserPermissionNotFound
	}
	return nil
}
