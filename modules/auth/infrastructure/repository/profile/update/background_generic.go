package update

import (
	"context"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	"nfxidentity/modules/auth/infrastructure/repository/profile/mapper"
)

//* =============================== Authority background_generic =============================== !//
func (h *Handler) AuthorityBackgroundGeneric(ctx context.Context, b *profileDomain.AuthorityProfileBackground) error {
	m := mapper.AuthorityProfileBackgroundDomainToModel(b)
	return h.db.WithContext(ctx).Save(m).Error
}

//* =============================== Forger background_generic =============================== !//
func (h *Handler) ForgerBackgroundGeneric(ctx context.Context, b *profileDomain.ForgerProfileBackground) error {
	m := mapper.ForgerProfileBackgroundDomainToModel(b)
	return h.db.WithContext(ctx).Save(m).Error
}
