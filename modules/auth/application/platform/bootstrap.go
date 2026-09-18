package platform

import (
	"context"
	"strings"
	"time"

	"nfxidentity/errors/src/auth"
	"nfxidentity/errors/src/sys"
	"nfxidentity/modules/auth/domain/account"
	"nfxidentity/modules/auth/domain/authorityprofile"
	"nfxidentity/modules/auth/domain/email"
	"nfxidentity/modules/auth/domain/forgerprofile"
	"nfxidentity/modules/auth/domain/identity"
	"nfxidentity/modules/auth/domain/phone"
	"nfxidentity/modules/auth/domain/settings"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
	"github.com/lib/pq"
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
	exists, err := s.repoFactory.Authority(none()).Get.AnyOwnerExists(ctx)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, auth.ErrAccountAlreadyExists
	}
	if _, err := s.repoFactory.Email(none()).Get.ByAddress(ctx, emailAddr); err == nil {
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
		forgerRepo := s.repoFactory.Forger(uow)
		authorityRepo := s.repoFactory.Authority(uow)
		settingsRepo := s.repoFactory.Settings(uow)
		if err := accountRepo.Create.New(ctx, account.NewFromState(account.AccountState{
			ID: accountID, AccountStatus: "active", SignupPlatform: "nfxidentity", CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if err := identityRepo.Create.New(ctx, identity.NewFromState(identity.IdentityState{
			ID: identityID, AccountID: accountID, IdentityProvider: "password", ProviderSubject: emailAddr,
			PasswordHash: &hashStr, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if err := emailRepo.Create.New(ctx, email.NewFromState(email.EmailState{
			ID: emailID, AccountID: accountID, Address: emailAddr, IsPrimary: true, VerifiedAt: &verified, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if phoneNum != "" {
			if err := phoneRepo.Create.New(ctx, phone.NewFromState(phone.PhoneState{
				ID: uuid.New(), AccountID: accountID, Number: phoneNum, IsPrimary: true, VerifiedAt: &verified, CreatedAt: now, UpdatedAt: now,
			})); err != nil {
				return err
			}
		}
		if err := forgerRepo.Create.New(ctx, forgerprofile.NewFromState(forgerprofile.State{
			ID: forgerID, AccountID: accountID, Roles: pq.StringArray{"forger"},
			ProfileLanguage: "zh", DisplayName: &display, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if err := settingsRepo.Create.New(ctx, settings.NewFromState(settings.State{
			ID: forgerID, Kind: "forger", LoginNotification: true, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if err := authorityRepo.Create.New(ctx, authorityprofile.NewFromState(authorityprofile.State{
			ID: authorityID, AccountID: accountID, Roles: pq.StringArray{"owner"},
			ProfileLanguage: "zh", DisplayName: &display, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		return settingsRepo.Create.New(ctx, settings.NewFromState(settings.State{
			ID: authorityID, Kind: "authority", LoginNotification: true, CreatedAt: now, UpdatedAt: now,
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
