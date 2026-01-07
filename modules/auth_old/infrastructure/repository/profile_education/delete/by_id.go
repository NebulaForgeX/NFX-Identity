package delete

import (
	"context"
	educationDomainErrors "nfxid/modules/auth/domain/profile_education/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 Education，实现 education.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Education{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return educationDomainErrors.ErrEducationNotFound
	}
	return nil
}
