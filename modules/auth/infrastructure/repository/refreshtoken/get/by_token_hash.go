package get

import (
	"context"
	"errors"

	"nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/refreshtoken/mapper"

	"gorm.io/gorm"
)

func (h *Handler) ByTokenHash(ctx context.Context, hash string) (*refreshtoken.RefreshToken, error) {
	var m models.Refreshtoken
	if err := h.db.WithContext(ctx).Where("token_hash = ? AND revoked_at IS NULL AND deleted_at IS NULL", hash).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, auth.ErrRefreshTokenNotFound
		}
		return nil, err
	}
	return mapper.ToDomain(&m), nil
}
