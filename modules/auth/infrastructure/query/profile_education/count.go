package profile_education

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// Count 获取 Education 总数，实现 education.Query 接口
func (h *Handler) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := h.db.WithContext(ctx).Model(&models.Education{}).Where("deleted_at IS NULL").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
