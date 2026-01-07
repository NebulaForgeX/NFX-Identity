package count

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// All 获取 ProfileBadge 总数，实现 profileBadgeDomain.Count 接口
func (h *Handler) All(ctx context.Context) (int64, error) {
	var count int64
	if err := h.db.WithContext(ctx).Model(&models.ProfileBadge{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
