package delete

import (
	"context"
	"nfxid/modules/auth/domain/password_history"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 PasswordHistory，实现 password_history.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.PasswordHistory{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return password_history.ErrPasswordHistoryNotFound
	}
	return nil
}
