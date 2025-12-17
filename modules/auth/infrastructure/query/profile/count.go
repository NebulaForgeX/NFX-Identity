package profile

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// Count 获取 Profile 总数，实现 profile.Query 接口
func (h *Handler) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := h.db.WithContext(ctx).Model(&models.Profile{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
