package delete

import (
	"context"
	"nfxid/modules/clients/domain/rate_limits"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/rate_limits/mapper"

	"github.com/google/uuid"
)

// ByAppIDAndLimitType 根据 AppID 和 LimitType 删除 RateLimit，实现 rate_limits.Delete 接口
func (h *Handler) ByAppIDAndLimitType(ctx context.Context, appID uuid.UUID, limitType rate_limits.RateLimitType) error {
	limitTypeEnum := mapper.RateLimitTypeDomainToEnum(limitType)
	result := h.db.WithContext(ctx).
		Where("app_id = ? AND limit_type = ?", appID, limitTypeEnum).
		Delete(&models.RateLimit{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return rate_limits.ErrRateLimitNotFound
	}
	return nil
}
