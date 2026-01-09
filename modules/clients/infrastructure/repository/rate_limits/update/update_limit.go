package update

import (
	"context"
	"time"
	"nfxid/modules/clients/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// UpdateLimit 更新 RateLimit 限制值，实现 rate_limits.Update 接口
func (h *Handler) UpdateLimit(ctx context.Context, id uuid.UUID, limitValue, windowSeconds int, description *string) error {
	updates := map[string]any{
		models.RateLimitCols.LimitValue:    limitValue,
		models.RateLimitCols.WindowSeconds: windowSeconds,
		models.RateLimitCols.Description:   description,
		models.RateLimitCols.UpdatedAt:     time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.RateLimit{}).
		Where("id = ?", id).
		Updates(updates).Error
}
