package delete

import (
	"context"
	"nfxid/modules/auth/domain/refresh_tokens"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 RefreshToken，实现 refresh_tokens.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.RefreshToken{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return refresh_tokens.ErrRefreshTokenNotFound
	}
	return nil
}
