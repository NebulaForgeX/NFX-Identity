package get

import (
	"context"
	"errors"

	authErr "nfxidentity/errors/src/auth"
	refreshtokenDomain "nfxidentity/modules/auth/domain/refreshtoken"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/refreshtoken/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*refreshtokenDomain.RefreshToken, error) {
	var m rdbmodels.RefreshToken
	if err := h.db.WithContext(ctx).Where(rdbmodels.RefreshTokenCols.ID+" = ?", id.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErr.ErrRefreshTokenNotFound
		}
		return nil, err
	}
	return mapper.RefreshTokenModelToDomain(&m), nil
}
