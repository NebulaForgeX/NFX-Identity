package profile_badge

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// Count 获取 ProfileBadge 总数，实现 profile_badge.Query 接口
func (h *Handler) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := h.db.WithContext(ctx).Model(&models.ProfileBadge{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
