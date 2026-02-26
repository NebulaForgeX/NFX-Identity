package get

import (
	"context"
	"errors"
	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/domain/refresh_tokens"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/refresh_tokens/mapper"

	"gorm.io/gorm"
)

// ByTokenID 根据 TokenID 获取 RefreshToken，实现 refresh_tokens.Get 接口
func (h *Handler) ByTokenID(ctx context.Context, tokenID string) (*refresh_tokens.RefreshToken, error) {
	var m models.RefreshToken
	if err := h.db.WithContext(ctx).Where("token_id = ?", tokenID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrRefreshTokenNotFound
		}
		return nil, err
	}
	return mapper.RefreshTokenModelToDomain(&m), nil
}
