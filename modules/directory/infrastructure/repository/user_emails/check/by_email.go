package check

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"
)

// ByEmail 根据 Email 检查 UserEmail 是否存在，实现 user_emails.Check 接口
func (h *Handler) ByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.UserEmail{}).
		Where("email = ?", email).
		Count(&count).Error
	return count > 0, err
}
