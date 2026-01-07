package check

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// ByUsername 根据 Username 检查 User 是否存在，实现 user.Check 接口
func (h *Handler) ByUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.User{}).
		Where("username = ?", username).
		Count(&count).Error
	return count > 0, err
}
