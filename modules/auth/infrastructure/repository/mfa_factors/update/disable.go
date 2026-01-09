package update

import (
	"context"
	"time"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// Disable 禁用 MFAFactor，实现 mfa_factors.Update 接口
func (h *Handler) Disable(ctx context.Context, factorID string) error {
	updates := map[string]any{
		models.MfaFactorCols.Enabled:   false,
		models.MfaFactorCols.UpdatedAt: time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.MfaFactor{}).
		Where("factor_id = ?", factorID).
		Updates(updates).Error
}
