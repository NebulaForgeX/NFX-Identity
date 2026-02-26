package delete

import (
	"context"
	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 Session，实现 sessions.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Session{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return authErr.ErrSessionNotFound
	}
	return nil
}
