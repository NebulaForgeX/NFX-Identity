package account

import (
	"context"
	"time"

	"github.com/google/uuid"

	"nfxidentity/constants"
	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	accountinternal "nfxidentity/modules/auth/application/account/internal"
	"nfxidentity/modules/auth/domain/profile"
	"nfxidentity/pkgs/transaction"
)

type CreateAuthorityProfileInput struct {
	AccountID       uuid.UUID
	DisplayName     string
	ProfileLanguage enums.AuthProfileLanguage
}

type CreateAuthorityProfileOutput struct {
	ProfileID uuid.UUID `json:"profile_id"`
}

func (s *Service) CreateAuthorityProfile(ctx context.Context, in CreateAuthorityProfileInput) (*CreateAuthorityProfileOutput, error) {
	if err := s.RequireOwner(ctx, in.AccountID); err != nil {
		return nil, err
	}
	authorityCount, err := s.repoFactory.Profile(accountinternal.None()).Check.CountAuthorityByAccountID(ctx, in.AccountID)
	if err != nil {
		return nil, authErr.ErrAuthorityProfileCountFailed.WithCause(err)
	}
	communityCount, err := s.repoFactory.Profile(accountinternal.None()).Check.CountForgerByAccountID(ctx, in.AccountID)
	if err != nil {
		return nil, authErr.ErrAuthorityProfileCountFailed.WithCause(err)
	}
	if authorityCount+communityCount >= constants.AuthAccountMaxProfiles {
		return nil, authErr.ErrAuthorityProfileLimitReached
	}

	now := time.Now()
	id := uuid.New()
	lang := enums.AuthProfileLanguage(accountinternal.LangOrDefault(string(in.ProfileLanguage)))
	err = s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		profileRepo := s.repoFactory.Profile(uow)
		if err := profileRepo.Create.NewAuthority(ctx, profile.NewAuthorityProfileFromState(profile.AuthorityProfileState{
			ID: id, AccountID: in.AccountID, AuthorityRoles: []enums.AuthAuthorityRole{enums.AuthAuthorityRoleAdministrator},
			ProfileLanguage: lang, Preference: profile.Default(lang), DisplayName: &in.DisplayName, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		return profileRepo.Create.NewAuthoritySettings(ctx, profile.NewAuthorityProfileSettingsFromState(profile.AuthorityProfileSettingsState{
			ID: id, LoginNotification: true, CreatedAt: now, UpdatedAt: now,
		}))
	})
	if err != nil {
		return nil, err
	}
	return &CreateAuthorityProfileOutput{ProfileID: id}, nil
}

type CreateCommunityProfileInput struct {
	AccountID       uuid.UUID
	DisplayName     string
	ProfileLanguage enums.AuthProfileLanguage
}

type CreateCommunityProfileOutput struct {
	ProfileID uuid.UUID `json:"profile_id"`
}

func (s *Service) CreateCommunityProfile(ctx context.Context, in CreateCommunityProfileInput) (*CreateCommunityProfileOutput, error) {
	communityCount, err := s.repoFactory.Profile(accountinternal.None()).Check.CountForgerByAccountID(ctx, in.AccountID)
	if err != nil {
		return nil, authErr.ErrForgerProfileCountFailed.WithCause(err)
	}
	authorityCount, err := s.repoFactory.Profile(accountinternal.None()).Check.CountAuthorityByAccountID(ctx, in.AccountID)
	if err != nil {
		return nil, authErr.ErrForgerProfileCountFailed.WithCause(err)
	}
	if communityCount+authorityCount >= constants.AuthAccountMaxProfiles {
		return nil, authErr.ErrForgerProfileLimitReached
	}

	now := time.Now()
	id := uuid.New()
	lang := enums.AuthProfileLanguage(accountinternal.LangOrDefault(string(in.ProfileLanguage)))
	err = s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		profileRepo := s.repoFactory.Profile(uow)
		if err := profileRepo.Create.NewForger(ctx, profile.NewForgerProfileFromState(profile.ForgerProfileState{
			ID: id, AccountID: in.AccountID, ForgerRoles: []enums.AuthForgerRole{enums.AuthForgerRoleForger},
			ProfileLanguage: lang, Preference: profile.Default(lang), DisplayName: &in.DisplayName, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		return profileRepo.Create.NewForgerSettings(ctx, profile.NewForgerProfileSettingsFromState(profile.ForgerProfileSettingsState{
			ID: id, LoginNotification: true, CreatedAt: now, UpdatedAt: now,
		}))
	})
	if err != nil {
		return nil, err
	}
	return &CreateCommunityProfileOutput{ProfileID: id}, nil
}
