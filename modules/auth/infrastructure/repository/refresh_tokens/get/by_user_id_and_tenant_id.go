package get

import (
	"context"
	"nfxid/modules/auth/domain/refresh_tokens"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/refresh_tokens/mapper"

	"github.com/google/uuid"
)

// ByUserIDAndTenantID 根据 UserID 和 TenantID 获取 RefreshToken 列表，实现 refresh_tokens.Get 接口
func (h *Handler) ByUserIDAndTenantID(ctx context.Context, userID, tenantID uuid.UUID) ([]*refresh_tokens.RefreshToken, error) {
	var ms []models.RefreshToken
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND tenant_id = ?", userID, tenantID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*refresh_tokens.RefreshToken, len(ms))
	for i, m := range ms {
		result[i] = mapper.RefreshTokenModelToDomain(&m)
	}
	return result, nil
}
