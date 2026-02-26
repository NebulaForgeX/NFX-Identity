package delete

import (
	"context"
	clientsErr "nfxidentity/errors/src/clients"
	"nfxidentity/modules/clients/infrastructure/rdb/models"

	"github.com/google/uuid"
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
		return clientsErr.ErrRateLimitNotFound
	}
	return nil
}
