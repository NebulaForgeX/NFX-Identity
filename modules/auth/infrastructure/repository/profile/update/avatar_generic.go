package update

import (
	"context"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	"nfxidentity/modules/auth/infrastructure/repository/profile/mapper"
)

//* =============================== Authority avatar_generic =============================== !//
func (h *Handler) AuthorityAvatarGeneric(ctx context.Context, a *profileDomain.AuthorityProfileAvatar) error {
	m := mapper.AuthorityProfileAvatarDomainToModel(a)
	return h.db.WithContext(ctx).Save(m).Error
}

//* =============================== Forger avatar_generic =============================== !//
func (h *Handler) ForgerAvatarGeneric(ctx context.Context, a *profileDomain.ForgerProfileAvatar) error {
	m := mapper.ForgerProfileAvatarDomainToModel(a)
	return h.db.WithContext(ctx).Save(m).Error
}
