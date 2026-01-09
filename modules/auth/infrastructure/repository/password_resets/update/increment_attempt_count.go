package update

import (
	"context"
	"time"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"gorm.io/gorm"
)

// IncrementAttemptCount 增加尝试次数，实现 password_resets.Update 接口
func (h *Handler) IncrementAttemptCount(ctx context.Context, resetID string) error {
	updates := map[string]any{
		models.PasswordResetCols.AttemptCount: gorm.Expr("attempt_count + 1"),
		models.PasswordResetCols.UpdatedAt:     time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.PasswordReset{}).
		Where("reset_id = ?", resetID).
		Updates(updates).Error
}
