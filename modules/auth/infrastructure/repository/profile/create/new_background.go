package create

import (
	"context"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	"nfxidentity/modules/auth/infrastructure/repository/profile/mapper"
)

//* =============================== Authority new_background =============================== !//
func (h *Handler) NewAuthorityBackground(ctx context.Context, b *profileDomain.AuthorityProfileBackground) error {
	return h.db.WithContext(ctx).Create(mapper.AuthorityProfileBackgroundDomainToModel(b)).Error
}

//* =============================== Forger new_background =============================== !//
func (h *Handler) NewForgerBackground(ctx context.Context, b *profileDomain.ForgerProfileBackground) error {
	return h.db.WithContext(ctx).Create(mapper.ForgerProfileBackgroundDomainToModel(b)).Error
}
