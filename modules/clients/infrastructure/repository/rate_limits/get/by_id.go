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

// ByID 根据 ID 获取 RateLimit，实现 rate_limits.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*rate_limits.RateLimit, error) {
	var m models.RateLimit
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, rate_limits.ErrRateLimitNotFound
		}
		return nil, err
	}
	return mapper.RateLimitModelToDomain(&m), nil
}
