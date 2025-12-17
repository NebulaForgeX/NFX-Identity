package delete

import (
	"context"
	occupationDomainErrors "nfxid/modules/auth/domain/profile_occupation/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 Occupation，实现 occupation.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Occupation{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return occupationDomainErrors.ErrOccupationNotFound
	}
	return nil
}
