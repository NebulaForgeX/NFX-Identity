package user

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// Count 获取 User 总数，实现 user.Query 接口
func (h *Handler) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := h.db.WithContext(ctx).Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
