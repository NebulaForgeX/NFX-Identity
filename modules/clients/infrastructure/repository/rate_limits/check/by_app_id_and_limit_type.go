package check

import (
	"context"
	"nfxid/modules/clients/domain/rate_limits"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/rate_limits/mapper"

	"github.com/google/uuid"
)

// ByAppIDAndLimitType 根据 AppID 和 LimitType 检查 RateLimit 是否存在，实现 rate_limits.Check 接口
func (h *Handler) ByAppIDAndLimitType(ctx context.Context, appID uuid.UUID, limitType rate_limits.RateLimitType) (bool, error) {
	limitTypeEnum := mapper.RateLimitTypeDomainToEnum(limitType)
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.RateLimit{}).
		Where("app_id = ? AND limit_type = ?", appID, limitTypeEnum).
		Count(&count).Error
	return count > 0, err
}
