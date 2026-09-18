package get

import (
	"context"
	"errors"

	authErr "nfxidentity/errors/src/auth"
	refreshtokenDomain "nfxidentity/modules/auth/domain/refreshtoken"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/refreshtoken/mapper"

	"gorm.io/gorm"
)

func (h *Handler) ByTokenHash(ctx context.Context, tokenHash string) (*refreshtokenDomain.RefreshToken, error) {
	var m rdbmodels.RefreshToken
	if err := h.db.WithContext(ctx).
		Where(rdbmodels.RefreshTokenCols.TokenHash+" = ?", tokenHash).
		Where(rdbmodels.RefreshTokenCols.RevokedAt + " IS NULL").
		Where(rdbmodels.RefreshTokenCols.DeletedAt + " IS NULL").
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrRefreshTokenNotFound
		}
		return nil, err
	}
	return mapper.RefreshTokenModelToDomain(&m), nil
}
