package account

import (
	"context"

	"github.com/google/uuid"

	"nfxidentity/constants"
	authErr "nfxidentity/errors/src/auth"
	accountinternal "nfxidentity/modules/auth/application/account/internal"
)

type DeleteAuthorityProfileInput struct {
	AccountID        uuid.UUID
	ProfileID        uuid.UUID
	CurrentProfileID uuid.UUID
}

func (s *Service) DeleteAuthorityProfile(ctx context.Context, in DeleteAuthorityProfileInput) error {
	if in.ProfileID == in.CurrentProfileID {
		return authErr.ErrAuthorityProfileCannotDeleteCurrent
	}
	profile, err := s.repoFactory.Profile(accountinternal.None()).Get.AuthorityByAccountAndID(ctx, in.AccountID, in.ProfileID)
	if err != nil {
		return authErr.ErrProfileNotOwned
	}
	authorityCount, err := s.repoFactory.Profile(accountinternal.None()).Check.CountAuthorityByAccountID(ctx, in.AccountID)
	if err != nil {
		return authErr.ErrAuthorityProfileCountFailed.WithCause(err)
	}
	communityCount, err := s.repoFactory.Profile(accountinternal.None()).Check.CountForgerByAccountID(ctx, in.AccountID)
	if err != nil {
		return authErr.ErrAuthorityProfileCountFailed.WithCause(err)
	}
	if authorityCount+communityCount <= constants.AuthAccountMinProfiles {
		return authErr.ErrAuthorityProfileCannotDeleteLast
	}
	if err := profile.Delete(); err != nil {
		return err
	}
	if err := s.repoFactory.Profile(accountinternal.None()).Update.AuthorityGeneric(ctx, profile); err != nil {
		return authErr.ErrAuthorityProfileDeleteFailed.WithCause(err)
	}
	return s.InvalidateFull(ctx, in.AccountID, in.ProfileID)
}

type DeleteCommunityProfileInput struct {
	AccountID        uuid.UUID
	ProfileID        uuid.UUID
	CurrentProfileID uuid.UUID
}

func (s *Service) DeleteCommunityProfile(ctx context.Context, in DeleteCommunityProfileInput) error {
	if in.ProfileID == in.CurrentProfileID {
		return authErr.ErrForgerProfileCannotDeleteCurrent
	}
	profile, err := s.repoFactory.Profile(accountinternal.None()).Get.ForgerByAccountAndID(ctx, in.AccountID, in.ProfileID)
	if err != nil {
		return authErr.ErrProfileNotOwned
	}
	communityCount, err := s.repoFactory.Profile(accountinternal.None()).Check.CountForgerByAccountID(ctx, in.AccountID)
	if err != nil {
		return authErr.ErrForgerProfileCountFailed.WithCause(err)
	}
	authorityCount, err := s.repoFactory.Profile(accountinternal.None()).Check.CountAuthorityByAccountID(ctx, in.AccountID)
	if err != nil {
		return authErr.ErrForgerProfileCountFailed.WithCause(err)
	}
	if communityCount+authorityCount <= constants.AuthAccountMinProfiles {
		return authErr.ErrForgerProfileCannotDeleteLast
	}
	if err := profile.Delete(); err != nil {
		return err
	}
	if err := s.repoFactory.Profile(accountinternal.None()).Update.ForgerGeneric(ctx, profile); err != nil {
		return authErr.ErrForgerProfileDeleteFailed.WithCause(err)
	}
	return s.InvalidateFull(ctx, in.AccountID, in.ProfileID)
}
