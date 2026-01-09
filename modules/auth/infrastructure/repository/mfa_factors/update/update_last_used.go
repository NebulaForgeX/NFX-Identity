package update

import (
	"context"
	"time"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// UpdateLastUsed 更新最后使用时间，实现 mfa_factors.Update 接口
func (h *Handler) UpdateLastUsed(ctx context.Context, factorID string) error {
	now := time.Now().UTC()
	updates := map[string]any{
		models.MfaFactorCols.LastUsedAt: &now,
		models.MfaFactorCols.UpdatedAt: now,
	}

	return h.db.WithContext(ctx).
		Model(&models.MfaFactor{}).
		Where("factor_id = ?", factorID).
		Updates(updates).Error
}
