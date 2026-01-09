package get

import (
	"context"
	"errors"
	"nfxid/modules/clients/domain/rate_limits"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/rate_limits/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByAppIDAndLimitType 根据 AppID 和 LimitType 获取 RateLimit，实现 rate_limits.Get 接口
func (h *Handler) ByAppIDAndLimitType(ctx context.Context, appID uuid.UUID, limitType rate_limits.RateLimitType) (*rate_limits.RateLimit, error) {
	limitTypeEnum := mapper.RateLimitTypeDomainToEnum(limitType)
	var m models.RateLimit
	if err := h.db.WithContext(ctx).
		Where("app_id = ? AND limit_type = ?", appID, limitTypeEnum).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, rate_limits.ErrRateLimitNotFound
		}
		return nil, err
	}
	return mapper.RateLimitModelToDomain(&m), nil
}
