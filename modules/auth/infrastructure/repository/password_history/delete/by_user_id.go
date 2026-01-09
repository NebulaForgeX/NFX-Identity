package delete

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserID 根据 UserID 删除 PasswordHistory，实现 password_history.Delete 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&models.PasswordHistory{}).Error
}
