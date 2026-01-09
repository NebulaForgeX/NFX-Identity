package create

import (
	"context"
	"nfxid/modules/clients/domain/rate_limits"
	"nfxid/modules/clients/infrastructure/repository/rate_limits/mapper"
)

// New 创建新的 RateLimit，实现 rate_limits.Create 接口
func (h *Handler) New(ctx context.Context, rl *rate_limits.RateLimit) error {
	m := mapper.RateLimitDomainToModel(rl)
	return h.db.WithContext(ctx).Create(&m).Error
}
