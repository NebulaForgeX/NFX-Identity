package check

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// ByFactorID 根据 FactorID 检查 MFAFactor 是否存在，实现 mfa_factors.Check 接口
func (h *Handler) ByFactorID(ctx context.Context, factorID string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.MfaFactor{}).
		Where("factor_id = ?", factorID).
		Count(&count).Error
	return count > 0, err
}
