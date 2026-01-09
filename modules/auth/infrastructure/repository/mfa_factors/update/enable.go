package update

import (
	"context"
	"time"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// Enable 启用 MFAFactor，实现 mfa_factors.Update 接口
func (h *Handler) Enable(ctx context.Context, factorID string) error {
	updates := map[string]any{
		models.MfaFactorCols.Enabled:   true,
		models.MfaFactorCols.UpdatedAt:  time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.MfaFactor{}).
		Where("factor_id = ?", factorID).
		Updates(updates).Error
}
