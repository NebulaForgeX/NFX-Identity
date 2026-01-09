package update

import (
	"context"
	"time"
	"nfxid/modules/auth/domain/password_resets"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/password_resets/mapper"
)

// UpdateStatus 更新状态，实现 password_resets.Update 接口
func (h *Handler) UpdateStatus(ctx context.Context, resetID string, status password_resets.ResetStatus) error {
	statusEnum := mapper.ResetStatusDomainToEnum(status)
	updates := map[string]any{
		models.PasswordResetCols.Status:    statusEnum,
		models.PasswordResetCols.UpdatedAt: time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.PasswordReset{}).
		Where("reset_id = ?", resetID).
		Updates(updates).Error
}
