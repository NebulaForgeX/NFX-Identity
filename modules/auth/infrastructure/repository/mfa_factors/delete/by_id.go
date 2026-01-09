package delete

import (
	"context"
	"nfxid/modules/auth/domain/mfa_factors"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 MFAFactor，实现 mfa_factors.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.MfaFactor{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return mfa_factors.ErrMFAFactorNotFound
	}
	return nil
}
