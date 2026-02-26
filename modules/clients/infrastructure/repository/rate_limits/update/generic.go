package update

import (
	"context"
	"nfxidentity/modules/clients/domain/rate_limits"
	"nfxidentity/modules/clients/infrastructure/rdb/models"
	"nfxidentity/modules/clients/infrastructure/repository/rate_limits/mapper"
)

// Generic 通用更新 RateLimit，实现 rate_limits.Update 接口
func (h *Handler) Generic(ctx context.Context, rl *rate_limits.RateLimit) error {
	m := mapper.RateLimitDomainToModel(rl)
	updates := mapper.RateLimitModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.RateLimit{}).
		Where("id = ?", rl.ID()).
		Updates(updates).Error
}
