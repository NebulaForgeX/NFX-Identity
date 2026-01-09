package update

import (
	"context"
	"time"
	"nfxid/modules/clients/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// UpdateStatus 更新 RateLimit 状态，实现 rate_limits.Update 接口
func (h *Handler) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	updates := map[string]any{
		models.RateLimitCols.Status:    status,
		models.RateLimitCols.UpdatedAt: time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.RateLimit{}).
		Where("id = ?", id).
		Updates(updates).Error
}
