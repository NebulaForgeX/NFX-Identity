package update

import (
	"context"
	"errors"
	"nfxid/enums"
	"nfxid/modules/auth/domain/password_resets"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"time"

	"gorm.io/gorm"
)

// MarkAsUsed 标记为已使用，实现 password_resets.Update 接口
func (h *Handler) MarkAsUsed(ctx context.Context, resetID string) error {
	// 先检查 PasswordReset 是否存在
	var m models.PasswordReset
	if err := h.db.WithContext(ctx).
		Where("reset_id = ?", resetID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return password_resets.ErrPasswordResetNotFound
		}
		return err
	}

	// 检查是否已经使用
	if m.UsedAt != nil {
		return password_resets.ErrResetAlreadyUsed
	}

	now := time.Now().UTC()
	status := enums.AuthResetStatusUsed
	updates := map[string]any{
		models.PasswordResetCols.UsedAt:    &now,
		models.PasswordResetCols.Status:    status,
		models.PasswordResetCols.UpdatedAt: now,
	}

	return h.db.WithContext(ctx).
		Model(&models.PasswordReset{}).
		Where("reset_id = ?", resetID).
		Updates(updates).Error
}
