package create

import (
	"context"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	"nfxidentity/modules/auth/infrastructure/repository/profile/mapper"
)

//* =============================== Authority new =============================== !//
func (h *Handler) NewAuthority(ctx context.Context, p *profileDomain.AuthorityProfile) error {
	return h.db.WithContext(ctx).Create(mapper.AuthorityProfileDomainToModel(p)).Error
}

//* =============================== Forger new =============================== !//
func (h *Handler) NewForger(ctx context.Context, p *profileDomain.ForgerProfile) error {
	return h.db.WithContext(ctx).Create(mapper.ForgerProfileDomainToModel(p)).Error
}
