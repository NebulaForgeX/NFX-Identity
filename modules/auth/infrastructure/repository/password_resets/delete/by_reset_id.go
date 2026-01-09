package delete

import (
	"context"
	"nfxid/modules/auth/domain/password_resets"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// ByResetID 根据 ResetID 删除 PasswordReset，实现 password_resets.Delete 接口
func (h *Handler) ByResetID(ctx context.Context, resetID string) error {
	result := h.db.WithContext(ctx).
		Where("reset_id = ?", resetID).
		Delete(&models.PasswordReset{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return password_resets.ErrPasswordResetNotFound
	}
	return nil
}
