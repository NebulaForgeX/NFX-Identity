package platform

import (
	"context"
	"nfxidentity/errors/src/auth"
	"nfxidentity/errors/src/sys"
	"strings"
	"time"

	"nfxidentity/modules/auth/domain/account"
	"nfxidentity/modules/auth/domain/email"
	"nfxidentity/modules/auth/domain/forgerprofile"
	"nfxidentity/modules/auth/domain/identity"
	"nfxidentity/modules/auth/domain/settings"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) SendSignupCode(ctx context.Context, emailAddr, lang string) error {
	emailAddr = strings.ToLower(strings.TrimSpace(emailAddr))
	if emailAddr == "" {
		return auth.ErrInvalidEmail
	}
	if _, err := s.repoFactory.Email(none()).Get.ByAddress(ctx, emailAddr); err == nil {
		return auth.ErrEmailAlreadyExists
	} else if !isMissing(err) {
		return auth.ErrEmailRegistrationCheckFailed.WithCause(err)
	}
	return s.issueAndStoreVerification(ctx, emailAddr, lang)
}

func (s *Service) SignupWithEmail(ctx context.Context, emailAddr, password, code, lang, deviceID, platformName string) (*LoginOutput, error) {
	emailAddr = strings.ToLower(strings.TrimSpace(emailAddr))
	if emailAddr == "" || password == "" {
		return nil, sys.ErrInvalidParams
	}
	if err := s.consumeVerificationCode(ctx, emailAddr, code); err != nil {
		return nil, err
	}
	if _, err := s.repoFactory.Email(none()).Get.ByAddress(ctx, emailAddr); err == nil {
		return nil, auth.ErrEmailAlreadyExists
	} else if !isMissing(err) {
		return nil, err
	}
	if platformName == "" {
		platformName = "nfxidentity"
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, auth.ErrHashFailed
	}
	now := time.Now()
	accountID := uuid.New()
	identityID := uuid.New()
	emailID := uuid.New()
	profileID := uuid.New()
	hashStr := string(hash)
	verified := now
	display := emailAddr

	err = s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		accountRepo := s.repoFactory.Account(uow)
		identityRepo := s.repoFactory.Identity(uow)
		emailRepo := s.repoFactory.Email(uow)
		forgerRepo := s.repoFactory.Forger(uow)
		settingsRepo := s.repoFactory.Settings(uow)
		if err := accountRepo.Create.New(ctx, account.NewFromState(account.AccountState{
			ID: accountID, AccountStatus: "active", SignupPlatform: platformName, CreatedAt: now, UpdatedAt: now,
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
		if err := forgerRepo.Create.New(ctx, forgerprofile.NewFromState(forgerprofile.State{
			ID: profileID, AccountID: accountID, Roles: pq.StringArray{"forger"},
			ProfileLanguage: langOrDefault(lang), DisplayName: &display, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		return settingsRepo.Create.New(ctx, settings.NewFromState(settings.State{
			ID: profileID, Kind: "forger", LoginNotification: true, CreatedAt: now, UpdatedAt: now,
		}))
	})
	if err != nil {
		return nil, auth.ErrSignupFailed.WithCause(err)
	}
	s.publishSignup(ctx, accountID, emailAddr, langOrDefault(lang))
	return s.issueAccountSession(ctx, accountID, &identityID, deviceID, emailAddr, "")
}

func (s *Service) LoginWithEmail(ctx context.Context, emailAddr, password, deviceID string) (*LoginOutput, error) {
	emailAddr = strings.ToLower(strings.TrimSpace(emailAddr))
	ident, err := s.repoFactory.Identity(none()).Get.ByProviderSubject(ctx, "password", emailAddr)
	if err != nil || ident.PasswordHash() == nil {
		return nil, auth.ErrInvalidCredentials
	}
	if bcrypt.CompareHashAndPassword([]byte(*ident.PasswordHash()), []byte(password)) != nil {
		return nil, auth.ErrInvalidCredentials
	}
	ident.TouchLogin(time.Now())
	_ = s.repoFactory.Identity(none()).Update.Generic(ctx, ident)
	return s.issueAccountSession(ctx, ident.AccountID(), ptrUUID(ident.ID()), deviceID, emailAddr, "")
}

func (s *Service) LoginWithPhone(ctx context.Context, phoneNum, password, deviceID string) (*LoginOutput, error) {
	phoneNum = strings.TrimSpace(phoneNum)
	p, err := s.repoFactory.Phone(none()).Get.ByNumber(ctx, phoneNum)
	if err != nil {
		return nil, auth.ErrInvalidCredentials
	}
	idents, err := s.repoFactory.Identity(none()).Get.ByAccountID(ctx, p.AccountID())
	if err != nil {
		return nil, auth.ErrInvalidCredentials
	}
	var ident *identity.Identity
	for _, item := range idents {
		if item.IdentityProvider() == "password" {
			ident = item
			break
		}
	}
	if ident == nil || ident.PasswordHash() == nil || bcrypt.CompareHashAndPassword([]byte(*ident.PasswordHash()), []byte(password)) != nil {
		return nil, auth.ErrInvalidCredentials
	}
	return s.issueAccountSession(ctx, p.AccountID(), ptrUUID(ident.ID()), deviceID, "", phoneNum)
}

func (s *Service) SelectProfile(ctx context.Context, accountID uuid.UUID, profileID uuid.UUID, kind, deviceID string) (*SelectProfileOutput, error) {
	emailAddr, phoneNum := s.primaryContacts(ctx, accountID)
	display := ""
	switch kind {
	case "forger":
		p, err := s.repoFactory.Forger(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return nil, auth.ErrProfileNotOwned
		}
		if p.DisplayName() != nil {
			display = *p.DisplayName()
		}
	case "authority":
		p, err := s.repoFactory.Authority(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return nil, auth.ErrProfileNotOwned
		}
		if p.DisplayName() != nil {
			display = *p.DisplayName()
		}
	default:
		return nil, auth.ErrInvalidProfileKind
	}
	access, refresh, err := s.tokens.GenerateTokenPair(accountID.String(), profileID.String(), display, emailAddr, phoneNum, kind)
	if err != nil {
		return nil, auth.ErrTokenFailed.WithCause(err)
	}
	if err := s.persistRefresh(ctx, accountID, nil, &profileID, &kind, deviceID, refresh); err != nil {
		return nil, err
	}
	provider, subject := s.latestIdentity(ctx, accountID)
	s.publishLogin(ctx, accountID, profileID, kind, provider, subject, emailAddr)
	return &SelectProfileOutput{
		AccountID: accountID.String(), ProfileID: profileID.String(), AccessToken: access, RefreshToken: refresh,
	}, nil
}

func (s *Service) Refresh(ctx context.Context, refreshToken, deviceID string) (*TokenOutput, error) {
	claims, err := s.tokens.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, auth.ErrInvalidRefreshToken
	}
	row, err := s.repoFactory.RefreshToken(none()).Get.ByTokenHash(ctx, hashToken(refreshToken))
	if err != nil {
		return nil, auth.ErrInvalidRefreshToken
	}
	now := time.Now()
	row.Revoke(now)
	_ = s.repoFactory.RefreshToken(none()).Update.Generic(ctx, row)
	access, refresh, err := s.tokens.GenerateTokenPair(claims.AccountID, claims.ProfileID, claims.Username, claims.Email, claims.Phone, claims.ProfileScope)
	if err != nil {
		return nil, auth.ErrTokenFailed.WithCause(err)
	}
	accountID, _ := uuid.Parse(claims.AccountID)
	var profileID *uuid.UUID
	if claims.ProfileID != "" && claims.ProfileID != emptyProfileID {
		id, err := uuid.Parse(claims.ProfileID)
		if err == nil {
			profileID = &id
		}
	}
	scope := claims.ProfileScope
	if err := s.persistRefresh(ctx, accountID, row.IdentityID(), profileID, nullableStr(scope), deviceID, refresh); err != nil {
		return nil, err
	}
	return &TokenOutput{AccessToken: access, RefreshToken: refresh}, nil
}

func ptrUUID(id uuid.UUID) *uuid.UUID { return &id }
