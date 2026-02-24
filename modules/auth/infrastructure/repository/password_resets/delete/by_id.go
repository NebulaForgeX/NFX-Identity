package delete

import (
	"context"
	authErr "nfxid/errors/src/auth"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 PasswordReset，实现 password_resets.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.PasswordReset{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return authErr.ErrPasswordResetNotFound
	}
	return nil
}
