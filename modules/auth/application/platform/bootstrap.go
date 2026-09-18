package platform

import (
	"context"
	"strings"
	"time"

	"nfxidentity/enums"
	"nfxidentity/errors/src/auth"
	"nfxidentity/errors/src/sys"
	"nfxidentity/modules/auth/domain/account"
	"nfxidentity/modules/auth/domain/email"
	"nfxidentity/modules/auth/domain/identity"
	"nfxidentity/modules/auth/domain/phone"
	"nfxidentity/modules/auth/domain/profile"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type BootstrapOwnerInput struct {
	Username string
	Password string
	Email    string
	Phone    string
}

type BootstrapOwnerResult struct {
	AccountID          string
	ForgerProfileID    string
	AuthorityProfileID string
}

func (s *Service) BootstrapOwner(ctx context.Context, in BootstrapOwnerInput) (*BootstrapOwnerResult, error) {
	username := strings.TrimSpace(in.Username)
	emailAddr := strings.ToLower(strings.TrimSpace(in.Email))
	phoneNum := strings.TrimSpace(in.Phone)
	if username == "" || in.Password == "" || emailAddr == "" {
		return nil, sys.ErrInvalidParams
	}
	exists, err := s.repoFactory.Profile(none()).Get.AnyOwnerExists(ctx)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, auth.ErrAccountAlreadyExists
	}
	if _, err := s.repoFactory.Email(none()).Get.ByEmail(ctx, emailAddr); err == nil {
		return nil, auth.ErrEmailAlreadyExists
	} else if !isMissing(err) {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, auth.ErrHashFailed
	}
	now := time.Now()
	accountID := uuid.New()
	identityID := uuid.New()
	emailID := uuid.New()
	forgerID := uuid.New()
	authorityID := uuid.New()
	hashStr := string(hash)
	verified := now
	display := username

	err = s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		accountRepo := s.repoFactory.Account(uow)
		identityRepo := s.repoFactory.Identity(uow)
		emailRepo := s.repoFactory.Email(uow)
		phoneRepo := s.repoFactory.Phone(uow)
		profileRepo := s.repoFactory.Profile(uow)

		if err := accountRepo.Create.New(ctx, account.NewAccountFromState(account.AccountState{
			ID: accountID, AccountStatus: enums.AuthAccountStatusActive, SignupPlatform: enums.AuthSignupPlatformNfxidentity, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if err := identityRepo.Create.New(ctx, identity.NewIdentityFromState(identity.IdentityState{
			ID: identityID, AccountID: accountID, IdentityProvider: enums.AuthIdentityProviderPassword, ProviderSubject: emailAddr,
			PasswordHash: &hashStr, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if err := emailRepo.Create.New(ctx, email.NewEmailFromState(email.EmailState{
			ID: emailID, AccountID: accountID, Email: emailAddr, IsPrimary: true, VerifiedAt: &verified, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if phoneNum != "" {
			if err := phoneRepo.Create.New(ctx, phone.NewPhoneFromState(phone.PhoneState{
				ID: uuid.New(), AccountID: accountID, Phone: phoneNum, IsPrimary: true, VerifiedAt: &verified, CreatedAt: now, UpdatedAt: now,
			})); err != nil {
				return err
			}
		}
		if err := profileRepo.Create.NewForger(ctx, profile.NewForgerProfileFromState(profile.ForgerProfileState{
			ID: forgerID, AccountID: accountID, ForgerRoles: []enums.AuthForgerRole{enums.AuthForgerRoleForger},
			ProfileLanguage: enums.AuthProfileLanguageZh, Preference: profile.Default(enums.AuthProfileLanguageZh), DisplayName: &display, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if err := profileRepo.Create.NewForgerSettings(ctx, profile.NewForgerProfileSettingsFromState(profile.ForgerProfileSettingsState{
			ID: forgerID, LoginNotification: true, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if err := profileRepo.Create.NewAuthority(ctx, profile.NewAuthorityProfileFromState(profile.AuthorityProfileState{
			ID: authorityID, AccountID: accountID, AuthorityRoles: []enums.AuthAuthorityRole{enums.AuthAuthorityRoleOwner},
			ProfileLanguage: enums.AuthProfileLanguageZh, Preference: profile.Default(enums.AuthProfileLanguageZh), DisplayName: &display, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		return profileRepo.Create.NewAuthoritySettings(ctx, profile.NewAuthorityProfileSettingsFromState(profile.AuthorityProfileSettingsState{
			ID: authorityID, LoginNotification: true, CreatedAt: now, UpdatedAt: now,
		}))
	})
	if err != nil {
		return nil, auth.ErrAccountCreateFailed.WithCause(err)
	}
	return &BootstrapOwnerResult{
		AccountID:          accountID.String(),
		ForgerProfileID:    forgerID.String(),
		AuthorityProfileID: authorityID.String(),
	}, nil
}

func (s *Service) Logout(ctx context.Context, refreshToken string) error {
	refreshToken = strings.TrimSpace(refreshToken)
	if refreshToken == "" {
		return auth.ErrInvalidRefreshToken
	}
	row, err := s.repoFactory.RefreshToken(none()).Get.ByTokenHash(ctx, hashToken(refreshToken))
	if err != nil {
		return auth.ErrInvalidRefreshToken
	}
	row.Revoke(time.Now())
	return s.repoFactory.RefreshToken(none()).Update.Generic(ctx, row)
}
