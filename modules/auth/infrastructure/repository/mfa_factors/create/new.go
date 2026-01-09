package create

import (
	"context"
	"nfxid/modules/auth/domain/mfa_factors"
	"nfxid/modules/auth/infrastructure/repository/mfa_factors/mapper"
)

// New 创建新的 MFAFactor，实现 mfa_factors.Create 接口
func (h *Handler) New(ctx context.Context, mf *mfa_factors.MFAFactor) error {
	m := mapper.MFAFactorDomainToModel(mf)
	return h.db.WithContext(ctx).Create(&m).Error
}
