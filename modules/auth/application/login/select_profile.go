package login

import (
	"context"
	"strings"
	"time"

	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/events"
	logininternal "nfxidentity/modules/auth/application/login/internal"
	accountdomain "nfxidentity/modules/auth/domain/account"
	emaildomain "nfxidentity/modules/auth/domain/email"
	refreshtokendomain "nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/pkgs/kafkax/eventbus"
	"nfxidentity/pkgs/logx"
	"nfxidentity/pkgs/ptrx"
	"nfxidentity/pkgs/tokenx"
	"nfxidentity/pkgs/tokenx/hashx"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
)

type SelectProfileInput struct {
	AccountID     uuid.UUID
	ProfileID     uuid.UUID
	Kind          enums.AuthProfileScope
	DeviceID      string
	CompleteLogin bool
	LoginEmail    string
}

// resolvedProfile is the thin, kind-agnostic projection the orchestration layer
// needs to mint a token. Each domain fills it from its own entity; there is no
// shared profile type between the two domains.
type resolvedProfile struct {
	profileID    uuid.UUID
	accountID    uuid.UUID
	displayName  *string
	profileScope enums.AuthProfileScope
}

type SelectProfileOutput struct {
	AccountID    uuid.UUID `json:"account_id"`
	ProfileID    uuid.UUID `json:"profile_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

// SelectProfile re-mints a token pair scoped to the chosen profile of the authenticated account.
func (s *Service) SelectProfile(ctx context.Context, in SelectProfileInput) (*SelectProfileOutput, error) {
	var account *accountdomain.Account
	var profile resolvedProfile
	var primaryEmail *emaildomain.Email
	var refreshTokenEntity *refreshtokendomain.RefreshToken
	var accessToken, refreshToken string
	var refreshClaims *tokenx.TokenClaims

	if err := s.txManager.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		accountRepo := s.repoFactory.Account(uow)
		emailRepo := s.repoFactory.Email(uow)
		refreshTokenRepo := s.repoFactory.RefreshToken(uow)
		var err error

		if account, err = accountRepo.Get.ByID(ctx, in.AccountID); err != nil {
			return authErr.ErrAccountNotFound
		}
		if !account.IsActive() {
			return authErr.ErrAccountInactive
		}

		//* =============================== Dispatch by kind (only routing, no domain logic) =============================== !//
		switch in.Kind {
		case enums.AuthProfileScopeCommunity:
			p, err := s.repoFactory.Profile(uow).Get.ForgerByProfileID(ctx, in.ProfileID)
			if err != nil {
				return authErr.ErrForgerProfileNotFound
			}
			if p.AccountID() != account.ID() {
				return authErr.ErrForgerProfileNotFound
			}
			profile = resolvedProfile{
				profileID:    p.ID(),
				accountID:    p.AccountID(),
				displayName:  p.DisplayName(),
				profileScope: enums.AuthProfileScopeCommunity,
			}
		case enums.AuthProfileScopeAuthority:
			p, err := s.repoFactory.Profile(uow).Get.AuthorityByProfileID(ctx, in.ProfileID)
			if err != nil {
				return authErr.ErrAuthorityProfileNotFound
			}
			if p.AccountID() != account.ID() {
				return authErr.ErrAuthorityProfileNotFound
			}
			profile = resolvedProfile{
				profileID:    p.ID(),
				accountID:    p.AccountID(),
				displayName:  p.DisplayName(),
				profileScope: enums.AuthProfileScopeAuthority,
			}
		default:
			return authErr.ErrForgerProfileNotFound
		}

		emails, err := emailRepo.Get.ByAccountID(ctx, account.ID())
		if err == nil {
			for _, e := range emails {
				if e.IsPrimary() {
					primaryEmail = e
					break
				}
			}
			if primaryEmail == nil && len(emails) > 0 {
				primaryEmail = emails[0]
			}
		}

		tokenEmail := ""
		if primaryEmail != nil {
			tokenEmail = primaryEmail.Email()
		}
		tokenUsername := tokenEmail
		if profile.displayName != nil && strings.TrimSpace(*profile.displayName) != "" {
			tokenUsername = strings.TrimSpace(*profile.displayName)
		}

		if accessToken, refreshToken, err = s.tokenx.GenerateTokenPair(
			account.ID().String(),
			profile.profileID.String(),
			tokenUsername,
			tokenEmail,
			"",
			string(profile.profileScope),
		); err != nil {
			return authErr.ErrLoginFailed.WithCause(err)
		}

		if refreshClaims, err = s.tokenx.VerifyRefreshToken(refreshToken); err != nil || refreshClaims == nil || refreshClaims.ExpiresAt == nil {
			return authErr.ErrLoginFailed.WithCause(err)
		}

		if refreshTokenEntity, err = refreshtokendomain.NewRefreshToken(refreshtokendomain.NewRefreshTokenParams{
			AccountID: account.ID(),
			DeviceID:  ptrx.Ptr(in.DeviceID),
			TokenHash: hashx.SHA256HexString(refreshToken),
			ExpiresAt: refreshClaims.ExpiresAt.Time,
		}); err != nil {
			return authErr.ErrLoginFailed.WithCause(err)
		} else if err := refreshTokenRepo.Create.New(ctx, refreshTokenEntity); err != nil {
			return authErr.ErrLoginFailed.WithCause(err)
		}
		if err := logininternal.RevokeOtherRefreshTokensForDevice(
			ctx,
			refreshTokenRepo,
			account.ID(),
			in.DeviceID,
			refreshTokenEntity.ID(),
			time.Now().UTC(),
		); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if in.CompleteLogin && s.bus != nil {
		loginAt := time.Now().UTC()
		providerSubject := ""
		if primaryEmail != nil {
			providerSubject = primaryEmail.Email()
		}
		if err := eventbus.PublishEvent(context.Background(), s.bus, events.LoginSuccessEvent{
			AccountID:        account.ID(),
			ProfileID:        profile.profileID,
			ProfileKind:      profile.profileScope,
			IdentityProvider: string(enums.AuthIdentityProviderPassword),
			ProviderSubject:  providerSubject,
			LoginEmail:       strings.TrimSpace(in.LoginEmail),
			LoginAt:          loginAt,
		}); err != nil {
			logx.S().Warnf("publish login success failed: %v", err)
		}
	}

	return &SelectProfileOutput{
		AccountID:    account.ID(),
		ProfileID:    profile.profileID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
