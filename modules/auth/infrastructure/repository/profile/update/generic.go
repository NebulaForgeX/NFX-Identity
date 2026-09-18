package update

import (
	"context"
	authErr "nfxidentity/errors/src/auth"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/profile/mapper"
)

//* =============================== Authority generic =============================== !//
func (h *Handler) AuthorityGeneric(ctx context.Context, p *profileDomain.AuthorityProfile) error {
	updates := mapper.AuthorityProfileDomainToUpdates(p)
	res := h.db.WithContext(ctx).Model(&rdbmodels.AuthorityProfile{}).
		Where(rdbmodels.AuthorityProfileCols.ID+" = ?", p.ID().String()).
		Updates(updates)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return authErr.ErrAuthorityProfileNotFound
	}
	return nil
}

//* =============================== Forger generic =============================== !//
func (h *Handler) ForgerGeneric(ctx context.Context, p *profileDomain.ForgerProfile) error {
	updates := mapper.ForgerProfileDomainToUpdates(p)
	res := h.db.WithContext(ctx).Model(&rdbmodels.ForgerProfile{}).
		Where(rdbmodels.ForgerProfileCols.ID+" = ?", p.ID().String()).
		Updates(updates)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return authErr.ErrForgerProfileNotFound
	}
	return nil
}
