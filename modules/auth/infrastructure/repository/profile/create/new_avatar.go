package create

import (
	"context"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	"nfxidentity/modules/auth/infrastructure/repository/profile/mapper"
)

//* =============================== Authority new_avatar =============================== !//
func (h *Handler) NewAuthorityAvatar(ctx context.Context, a *profileDomain.AuthorityProfileAvatar) error {
	return h.db.WithContext(ctx).Create(mapper.AuthorityProfileAvatarDomainToModel(a)).Error
}

//* =============================== Forger new_avatar =============================== !//
func (h *Handler) NewForgerAvatar(ctx context.Context, a *profileDomain.ForgerProfileAvatar) error {
	return h.db.WithContext(ctx).Create(mapper.ForgerProfileAvatarDomainToModel(a)).Error
}
