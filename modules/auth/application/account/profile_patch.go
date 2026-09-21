package account

import (
	"context"

	"github.com/google/uuid"

	authErr "nfxidentity/errors/src/auth"
	accountinternal "nfxidentity/modules/auth/application/account/internal"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	"nfxidentity/pkgs/patch"
)

type PatchAuthorityProfileInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	Patch     profileDomain.AuthorityProfilePatch
}

func (s *Service) PatchAuthorityProfile(ctx context.Context, in PatchAuthorityProfileInput) error {
	profile, err := s.repoFactory.Profile(accountinternal.None()).Get.AuthorityByAccountAndID(ctx, in.AccountID, in.ProfileID)
	if err != nil {
		return authErr.ErrProfileNotOwned
	}

	minimal := in.Patch.DiffAgainst(profile)
	if patch.IsPatchEmpty(minimal) {
		return nil
	} else if err := profile.ApplyPatch(minimal); err != nil {
		return err
	}

	if err := s.repoFactory.Profile(accountinternal.None()).Update.AuthorityPartial(ctx, in.ProfileID, minimal); err != nil {
		return authErr.ErrAuthorityProfileUpdateFailed.WithCause(err)
	}
	return s.InvalidateFull(ctx, in.AccountID, in.ProfileID)
}

type PatchCommunityProfileInput struct {
	AccountID uuid.UUID
	ProfileID uuid.UUID
	Patch     profileDomain.ForgerProfilePatch
}

func (s *Service) PatchCommunityProfile(ctx context.Context, in PatchCommunityProfileInput) error {
	profile, err := s.repoFactory.Profile(accountinternal.None()).Get.ForgerByAccountAndID(ctx, in.AccountID, in.ProfileID)
	if err != nil {
		return authErr.ErrProfileNotOwned
	}

	minimal := in.Patch.DiffAgainst(profile)
	if patch.IsPatchEmpty(minimal) {
		return nil
	} else if err := profile.ApplyPatch(minimal); err != nil {
		return err
	}

	if err := s.repoFactory.Profile(accountinternal.None()).Update.ForgerPartial(ctx, in.ProfileID, minimal); err != nil {
		return authErr.ErrForgerProfileUpdateFailed.WithCause(err)
	}
	return s.InvalidateFull(ctx, in.AccountID, in.ProfileID)
}
