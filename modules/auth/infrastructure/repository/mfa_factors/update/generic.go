package update

import (
	"context"
	"nfxid/modules/auth/domain/mfa_factors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mfa_factors/mapper"
)

// Generic 通用更新 MFAFactor，实现 mfa_factors.Update 接口
func (h *Handler) Generic(ctx context.Context, mf *mfa_factors.MFAFactor) error {
	m := mapper.MFAFactorDomainToModel(mf)
	updates := mapper.MFAFactorModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.MfaFactor{}).
		Where("id = ?", mf.ID()).
		Updates(updates).Error
}
