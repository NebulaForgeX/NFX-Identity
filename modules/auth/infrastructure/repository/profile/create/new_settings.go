package create

import (
	"context"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	"nfxidentity/modules/auth/infrastructure/repository/profile/mapper"
)

//* =============================== Authority new_settings =============================== !//
func (h *Handler) NewAuthoritySettings(ctx context.Context, s *profileDomain.AuthorityProfileSettings) error {
	return h.db.WithContext(ctx).Create(mapper.AuthorityProfileSettingsDomainToModel(s)).Error
}

//* =============================== Forger new_settings =============================== !//
func (h *Handler) NewForgerSettings(ctx context.Context, s *profileDomain.ForgerProfileSettings) error {
	return h.db.WithContext(ctx).Create(mapper.ForgerProfileSettingsDomainToModel(s)).Error
}
