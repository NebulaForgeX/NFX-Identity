package check

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// ByPhone 根据 Phone 检查 User 是否存在，实现 user.Check 接口
func (h *Handler) ByPhone(ctx context.Context, phone string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.User{}).
		Where("phone = ?", phone).
		Count(&count).Error
	return count > 0, err
}
