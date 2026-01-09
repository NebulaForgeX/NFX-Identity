package delete

import (
	"context"
	"nfxid/modules/auth/domain/mfa_factors"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// ByFactorID 根据 FactorID 删除 MFAFactor，实现 mfa_factors.Delete 接口
func (h *Handler) ByFactorID(ctx context.Context, factorID string) error {
	result := h.db.WithContext(ctx).
		Where("factor_id = ?", factorID).
		Delete(&models.MfaFactor{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return mfa_factors.ErrMFAFactorNotFound
	}
	return nil
}
