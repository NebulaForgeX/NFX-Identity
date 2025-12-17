package delete

import (
	"context"
	userRoleErrors "nfxid/modules/auth/domain/user_role/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 UserRole，实现 user_role.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.UserRole{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return userRoleErrors.ErrUserRoleNotFound
	}
	return nil
}
