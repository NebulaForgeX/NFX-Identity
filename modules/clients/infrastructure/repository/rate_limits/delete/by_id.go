package delete

import (
	"context"
	"errors"
	"nfxid/modules/clients/domain/rate_limits"
	"nfxid/modules/clients/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 删除 RateLimit，实现 rate_limits.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.RateLimit{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return rate_limits.ErrRateLimitNotFound
	}
	return nil
}
