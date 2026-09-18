package update

import (
	"context"

	authErr "nfxidentity/errors/src/auth"
	identityDomain "nfxidentity/modules/auth/domain/identity"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/identity/mapper"
)

func (h *Handler) Generic(ctx context.Context, i *identityDomain.Identity) error {
	updates := mapper.IdentityDomainToUpdates(i)
	res := h.db.WithContext(ctx).Model(&rdbmodels.Identity{}).
		Where(rdbmodels.IdentityCols.ID+" = ?", i.ID().String()).
		Updates(updates)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return authErr.ErrIdentityNotFound
	}
	return nil
}
