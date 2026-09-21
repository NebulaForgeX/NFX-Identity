package signup

import (
	"context"
	"strings"
	"time"

	"nfxidentity/constants"
	"nfxidentity/enums"
	"nfxidentity/events"
	authErr "nfxidentity/errors/src/auth"
	signupinternal "nfxidentity/modules/auth/application/signup/internal"
	accountdomain "nfxidentity/modules/auth/domain/account"
	emaildomain "nfxidentity/modules/auth/domain/email"
	identitydomain "nfxidentity/modules/auth/domain/identity"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	refreshtokendomain "nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/pkgs/kafkax/eventbus"
	"nfxidentity/pkgs/logx"
	"nfxidentity/pkgs/ptrx"
	"nfxidentity/pkgs/tokenx"
	"nfxidentity/pkgs/tokenx/hashx"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type WithEmailInput struct {
	Email            string
	Password         string
	VerificationCode string
	Lang             enums.AuthProfileLanguage
	DeviceID         string
	SignupPlatform   enums.AuthSignupPlatform
}

type WithEmailOutput struct {
	AccountID    uuid.UUID `json:"account_id"`
	ProfileID    uuid.UUID `json:"profile_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

func (s *Service) WithEmail(ctx context.Context, in WithEmailInput) (*WithEmailOutput, error) {
	cachedCode, err := s.lookupVerificationCode(ctx, in.Email)
	if err != nil {
		return nil, err
	} else if !strings.EqualFold(strings.TrimSpace(cachedCode), in.VerificationCode) {
		return nil, authErr.ErrVerificationCodeWrong
	}

	in.Lang = constants.AuthLanguage.Parse(in.Lang, enums.AuthProfileLanguageEn)
	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, authErr.ErrSignupFailed.WithCause(err)
	}

	var account *accountdomain.Account
	var profile *profileDomain.ForgerProfile
	var email *emaildomain.Email
	var identity *identitydomain.Identity
	var accessToken string
	var refreshToken string
	var refreshClaims *tokenx.TokenClaims
	var refreshTokenEntity *refreshtokendomain.RefreshToken
	if err := s.tx.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		accountRepo := s.repoFactory.Account(uow)
		emailRepo := s.repoFactory.Email(uow)
		profileRepo := s.repoFactory.Profile(uow)
		identityRepo := s.repoFactory.Identity(uow)
		refreshTokenRepo := s.repoFactory.RefreshToken(uow)
		var err error
		var registered bool

		if registered, err = emailRepo.Check.ByEmail(ctx, in.Email); err != nil {
			return authErr.ErrEmailRegistrationCheckFailed.WithCause(err)
		} else if registered {
			return authErr.ErrEmailAlreadyExists
		}

		if account, err = accountdomain.NewAccount(accountdomain.NewAccountParams{
			AccountStatus:  enums.AuthAccountStatusActive,
			SignupPlatform: in.SignupPlatform,
		}); err != nil {
			return authErr.ErrSignupFailed.WithCause(err)
		} else if err := accountRepo.Create.New(ctx, account); err != nil {
			return authErr.ErrAccountCreateFailed.WithCause(err)
		}

		if email, err = emaildomain.NewEmail(emaildomain.NewEmailParams{
			AccountID: account.ID(),
			Email:     in.Email,
			IsPrimary: true,
		}); err != nil {
			return authErr.ErrSignupFailed.WithCause(err)
		} else if err := email.MarkVerified(time.Now().UTC()); err != nil {
			return authErr.ErrSignupFailed.WithCause(err)
		} else if err := emailRepo.Create.New(ctx, email); err != nil {
			return authErr.ErrAccountCreateFailed.WithCause(err)
		}

		if profile, err = profileDomain.NewForgerProfile(profileDomain.NewForgerProfileParams{
			AccountID:       account.ID(),
			ProfileLanguage: in.Lang,
			Preference:      profileDomain.Default(in.Lang),
			DisplayName:     ptrx.Ptr(signupinternal.GenerateForgerDisplayName(in.Email)),
			Timezone:        ptrx.Ptr("UTC"),
		}); err != nil {
			return authErr.ErrSignupFailed.WithCause(err)
		} else if err := profileRepo.Create.NewForger(ctx, profile); err != nil {
			return authErr.ErrAccountCreateFailed.WithCause(err)
		}

		var settings *profileDomain.ForgerProfileSettings
		if settings, err = profileDomain.NewForgerProfileSettings(profileDomain.NewForgerProfileSettingsParams{
			ProfileID:         profile.ID(),
			LoginNotification: true,
		}); err != nil {
			return authErr.ErrSignupFailed.WithCause(err)
		} else if err := profileRepo.Create.NewForgerSettings(ctx, settings); err != nil {
			return authErr.ErrForgerProfileSettingsCreateFailed.WithCause(err)
		}

		if accessToken, refreshToken, err = s.tokens.GenerateTokenPair(
			account.ID().String(),
			profile.ID().String(),
			*profile.DisplayName(),
			email.Email(),
			"",
			string(enums.AuthProfileScopeCommunity),
		); err != nil {
			return authErr.ErrSignupFailed.WithCause(err)
		}

		if refreshClaims, err = s.tokens.VerifyRefreshToken(refreshToken); err != nil || refreshClaims == nil || refreshClaims.ExpiresAt == nil {
			return authErr.ErrSignupFailed.WithCause(err)
		}

		if identity, err = identitydomain.NewIdentity(identitydomain.NewIdentityParams{
			AccountID:        account.ID(),
			IdentityProvider: enums.AuthIdentityProviderPassword,
			ProviderSubject:  in.Email,
			PasswordHash:     ptrx.Ptr(string(passwordHashBytes)),
		}); err != nil {
			return authErr.ErrSignupFailed.WithCause(err)
		} else if err := identityRepo.Create.New(ctx, identity); err != nil {
			return authErr.ErrSignupFailed.WithCause(err)
		}

		if refreshTokenEntity, err = refreshtokendomain.NewRefreshToken(refreshtokendomain.NewRefreshTokenParams{
			AccountID:  account.ID(),
			IdentityID: ptrx.Ptr(identity.ID()),
			DeviceID:   ptrx.Ptr(in.DeviceID),
			TokenHash:  hashx.SHA256HexString(refreshToken),
			ExpiresAt:  refreshClaims.ExpiresAt.Time,
		}); err != nil {
			return authErr.ErrSignupFailed.WithCause(err)
		} else if err := refreshTokenRepo.Create.New(ctx, refreshTokenEntity); err != nil {
			return authErr.ErrSignupFailed.WithCause(err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	if s.bus != nil {
		if err := eventbus.PublishEvent(context.Background(), s.bus, events.SignupSuccessEvent{
			AccountID: account.ID(),
			Email:     in.Email,
			Lang:      string(in.Lang),
		}); err != nil {
			logx.L().Warn("signup: failed to publish signup success", zap.Error(err))
		}
	}

	if err := s.deleteVerificationCode(ctx, in.Email); err != nil {
		logx.L().Warn("signup: failed to delete verification code cache",
			zap.String("email", in.Email),
			zap.Error(err),
		)
	}

	return &WithEmailOutput{
		AccountID:    account.ID(),
		ProfileID:    profile.ID(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
