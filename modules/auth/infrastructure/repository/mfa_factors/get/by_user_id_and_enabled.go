package get

import (
	"context"
	"nfxid/modules/auth/domain/mfa_factors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mfa_factors/mapper"

	"github.com/google/uuid"
)

// ByUserIDAndEnabled 根据 UserID 和 Enabled 获取 MFAFactor 列表，实现 mfa_factors.Get 接口
func (h *Handler) ByUserIDAndEnabled(ctx context.Context, userID uuid.UUID, enabled bool) ([]*mfa_factors.MFAFactor, error) {
	var ms []models.MfaFactor
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND enabled = ?", userID, enabled).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*mfa_factors.MFAFactor, len(ms))
	for i, m := range ms {
		result[i] = mapper.MFAFactorModelToDomain(&m)
	}
	return result, nil
}
