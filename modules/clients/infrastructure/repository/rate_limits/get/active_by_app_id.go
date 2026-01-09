package get

import (
	"context"
	"nfxid/modules/clients/domain/rate_limits"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/rate_limits/mapper"

	"github.com/google/uuid"
)

// ActiveByAppID 根据 AppID 获取活跃的 RateLimit 列表，实现 rate_limits.Get 接口
func (h *Handler) ActiveByAppID(ctx context.Context, appID uuid.UUID) ([]*rate_limits.RateLimit, error) {
	var ms []models.RateLimit
	if err := h.db.WithContext(ctx).
		Where("app_id = ? AND status = ?", appID, "active").
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*rate_limits.RateLimit, len(ms))
	for i, m := range ms {
		result[i] = mapper.RateLimitModelToDomain(&m)
	}
	return result, nil
}
