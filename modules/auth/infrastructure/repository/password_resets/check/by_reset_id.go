package check

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// ByResetID 根据 ResetID 检查 PasswordReset 是否存在，实现 password_resets.Check 接口
func (h *Handler) ByResetID(ctx context.Context, resetID string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.PasswordReset{}).
		Where("reset_id = ?", resetID).
		Count(&count).Error
	return count > 0, err
}
