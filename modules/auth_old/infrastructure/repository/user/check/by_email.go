package check

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// ByEmail 根据 Email 检查 User 是否存在，实现 user.Check 接口
func (h *Handler) ByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.User{}).
		Where("email = ?", email).
		Count(&count).Error
	return count > 0, err
}
