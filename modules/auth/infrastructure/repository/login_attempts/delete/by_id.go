package delete

import (
	"context"
	"nfxid/modules/auth/domain/login_attempts"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 LoginAttempt，实现 login_attempts.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.LoginAttempt{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return login_attempts.ErrLoginAttemptNotFound
	}
	return nil
}
