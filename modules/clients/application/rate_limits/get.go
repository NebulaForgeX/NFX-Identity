package rate_limits

import (
	"context"
	rateLimitResult "nfxid/modules/clients/application/rate_limits/results"

	"github.com/google/uuid"
)

// GetRateLimit 根据ID获取速率限制
func (s *Service) GetRateLimit(ctx context.Context, rateLimitID uuid.UUID) (rateLimitResult.RateLimitRO, error) {
	domainEntity, err := s.rateLimitRepo.Get.ByID(ctx, rateLimitID)
	if err != nil {
		return rateLimitResult.RateLimitRO{}, err
	}
	return rateLimitResult.RateLimitMapper(domainEntity), nil
}
