package platform

import (
	"context"
	"strings"
	"time"

	"nfxidentity/modules/auth/domain/account"
	"nfxidentity/modules/auth/domain/email"
	"nfxidentity/modules/auth/domain/forgerprofile"
	"nfxidentity/modules/auth/domain/identity"
	"nfxidentity/modules/auth/domain/settings"
	"nfxidentity/pkgs/errx"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) SendSignupCode(ctx context.Context, emailAddr string) error {
	emailAddr = strings.ToLower(strings.TrimSpace(emailAddr))
	if emailAddr == "" {
		return errx.InvalidArg("INVALID_EMAIL", "email required")
	}
	s.StoreVerificationCode(ctx, emailAddr, RandomCode())
	return nil
}

func (s *Service) SignupWithEmail(ctx context.Context, emailAddr, password, code, lang, deviceID, platformName string) (*LoginOutput, error) {
	emailAddr = strings.ToLower(strings.TrimSpace(emailAddr))
	if emailAddr == "" || password == "" {
		return nil, errx.InvalidArg("INVALID_PARAMS", "email and password required")
	}
	if !s.checkVerificationCode(ctx, emailAddr, code) {
		return nil, errx.InvalidArg("INVALID_VERIFICATION_CODE", "invalid verification code")
	}
	if platformName == "" {
		platformName = "nfxidentity"
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errx.Internal("HASH_FAILED", "failed to hash password")
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
		if err := s.repos.Account(uow).Create.New(ctx, account.NewFromState(account.AccountState{
			ID: accountID, AccountStatus: "active", SignupPlatform: platformName, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if err := s.repos.Identity(uow).Create.New(ctx, identity.NewFromState(identity.IdentityState{
			ID: identityID, AccountID: accountID, IdentityProvider: "password", ProviderSubject: emailAddr,
			PasswordHash: &hashStr, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if err := s.repos.Email(uow).Create.New(ctx, email.NewFromState(email.EmailState{
			ID: emailID, AccountID: accountID, Address: emailAddr, IsPrimary: true, VerifiedAt: &verified, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		if err := s.repos.Forger(uow).Create.New(ctx, forgerprofile.NewFromState(forgerprofile.State{
			ID: profileID, AccountID: accountID, Roles: pq.StringArray{"forger"},
			ProfileLanguage: langOrDefault(lang), DisplayName: &display, CreatedAt: now, UpdatedAt: now,
		})); err != nil {
			return err
		}
		return s.repos.Settings(uow).Create.New(ctx, settings.NewFromState(settings.State{
			ID: profileID, Kind: "forger", LoginNotification: true, CreatedAt: now, UpdatedAt: now,
		}))
	})
	if err != nil {
		return nil, errx.Internal("SIGNUP_FAILED", err.Error())
	}
	return s.issueAccountSession(ctx, accountID, &identityID, deviceID, emailAddr, "")
}

func (s *Service) LoginWithEmail(ctx context.Context, emailAddr, password, deviceID string) (*LoginOutput, error) {
	emailAddr = strings.ToLower(strings.TrimSpace(emailAddr))
	ident, err := s.repos.Identity(none()).Get.ByProviderSubject(ctx, "password", emailAddr)
	if err != nil || ident.PasswordHash() == nil {
		return nil, ErrInvalidCredentials
	}
	if bcrypt.CompareHashAndPassword([]byte(*ident.PasswordHash()), []byte(password)) != nil {
		return nil, ErrInvalidCredentials
	}
	ident.TouchLogin(time.Now())
	_ = s.repos.Identity(none()).Update.Generic(ctx, ident)
	return s.issueAccountSession(ctx, ident.AccountID(), ptrUUID(ident.ID()), deviceID, emailAddr, "")
}

func (s *Service) LoginWithPhone(ctx context.Context, phoneNum, password, deviceID string) (*LoginOutput, error) {
	phoneNum = strings.TrimSpace(phoneNum)
	p, err := s.repos.Phone(none()).Get.ByNumber(ctx, phoneNum)
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	idents, err := s.repos.Identity(none()).Get.ByAccountID(ctx, p.AccountID())
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	var ident *identity.Identity
	for _, item := range idents {
		if item.IdentityProvider() == "password" {
			ident = item
			break
		}
	}
	if ident == nil || ident.PasswordHash() == nil || bcrypt.CompareHashAndPassword([]byte(*ident.PasswordHash()), []byte(password)) != nil {
		return nil, ErrInvalidCredentials
	}
	return s.issueAccountSession(ctx, p.AccountID(), ptrUUID(ident.ID()), deviceID, "", phoneNum)
}

func (s *Service) SelectProfile(ctx context.Context, accountID uuid.UUID, profileID uuid.UUID, kind, deviceID string) (*SelectProfileOutput, error) {
	emailAddr, phoneNum := s.primaryContacts(ctx, accountID)
	display := ""
	switch kind {
	case "forger":
		p, err := s.repos.Forger(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return nil, ErrProfileNotOwned
		}
		if p.DisplayName() != nil {
			display = *p.DisplayName()
		}
	case "authority":
		p, err := s.repos.Authority(none()).Get.ByAccountAndID(ctx, accountID, profileID)
		if err != nil {
			return nil, ErrProfileNotOwned
		}
		if p.DisplayName() != nil {
			display = *p.DisplayName()
		}
	default:
		return nil, errx.InvalidArg("INVALID_PROFILE_KIND", "kind must be forger or authority")
	}
	access, refresh, err := s.tokens.GenerateTokenPair(accountID.String(), profileID.String(), display, emailAddr, phoneNum, kind)
	if err != nil {
		return nil, errx.Internal("TOKEN_FAILED", err.Error())
	}
	if err := s.persistRefresh(ctx, accountID, nil, &profileID, &kind, deviceID, refresh); err != nil {
		return nil, err
	}
	return &SelectProfileOutput{
		AccountID: accountID.String(), ProfileID: profileID.String(), AccessToken: access, RefreshToken: refresh,
	}, nil
}

func (s *Service) Refresh(ctx context.Context, refreshToken, deviceID string) (*TokenOutput, error) {
	claims, err := s.tokens.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, ErrInvalidRefresh
	}
	row, err := s.repos.RefreshToken(none()).Get.ByTokenHash(ctx, hashToken(refreshToken))
	if err != nil {
		return nil, ErrInvalidRefresh
	}
	now := time.Now()
	row.Revoke(now)
	_ = s.repos.RefreshToken(none()).Update.Generic(ctx, row)
	access, refresh, err := s.tokens.GenerateTokenPair(claims.AccountID, claims.ProfileID, claims.Username, claims.Email, claims.Phone, claims.ProfileScope)
	if err != nil {
		return nil, errx.Internal("TOKEN_FAILED", err.Error())
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
