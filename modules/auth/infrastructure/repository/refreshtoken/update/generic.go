package update

import (
	"context"

	authErr "nfxidentity/errors/src/auth"
	refreshtokenDomain "nfxidentity/modules/auth/domain/refreshtoken"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/refreshtoken/mapper"
)

func (h *Handler) Generic(ctx context.Context, t *refreshtokenDomain.RefreshToken) error {
	updates := mapper.RefreshTokenDomainToUpdates(t)
	res := h.db.WithContext(ctx).Model(&rdbmodels.RefreshToken{}).
		Where(rdbmodels.RefreshTokenCols.ID+" = ?", t.ID().String()).
		Updates(updates)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return authErr.ErrRefreshTokenNotFound
	}
	return nil
}
