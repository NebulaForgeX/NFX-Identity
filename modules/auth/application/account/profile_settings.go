package account

import (
	"context"

	"github.com/google/uuid"

	"nfxidentity/enums"
	accountinternal "nfxidentity/modules/auth/application/account/internal"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	"nfxidentity/pkgs/patch"
)

type PatchAuthorityProfileSettingsInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	Patch     profileDomain.AuthorityProfileSettingsPatch
}

func (s *Service) PatchAuthorityProfileSettings(ctx context.Context, in PatchAuthorityProfileSettingsInput) error {
	if patch.IsPatchEmpty(in.Patch) {
		return nil
	}
	if err := s.requireOwnedProfile(ctx, in.AccountID, in.ProfileID, enums.AuthProfileScopeAuthority); err != nil {
		return err
	}
	return s.repoFactory.Profile(accountinternal.None()).Update.AuthorityPartialSettings(ctx, in.ProfileID, in.Patch)
}

type PatchCommunityProfileSettingsInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	Patch     profileDomain.ForgerProfileSettingsPatch
}

func (s *Service) PatchCommunityProfileSettings(ctx context.Context, in PatchCommunityProfileSettingsInput) error {
	if patch.IsPatchEmpty(in.Patch) {
		return nil
	}
	if err := s.requireOwnedProfile(ctx, in.AccountID, in.ProfileID, enums.AuthProfileScopeCommunity); err != nil {
		return err
	}
	return s.repoFactory.Profile(accountinternal.None()).Update.ForgerPartialSettings(ctx, in.ProfileID, in.Patch)
}
