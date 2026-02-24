package get

import (
	"context"
	"errors"
	authErr "nfxid/errors/src/auth"
	"nfxid/modules/auth/domain/mfa_factors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mfa_factors/mapper"

	"gorm.io/gorm"
)

// ByFactorID 根据 FactorID 获取 MFAFactor，实现 mfa_factors.Get 接口
func (h *Handler) ByFactorID(ctx context.Context, factorID string) (*mfa_factors.MFAFactor, error) {
	var m models.MfaFactor
	if err := h.db.WithContext(ctx).Where("factor_id = ?", factorID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrMFAFactorNotFound
		}
		return nil, err
	}
	return mapper.MFAFactorModelToDomain(&m), nil
}
