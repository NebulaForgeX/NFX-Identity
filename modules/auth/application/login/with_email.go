package login

import (
	"context"
	"time"

	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	logininternal "nfxidentity/modules/auth/application/login/internal"
	accountdomain "nfxidentity/modules/auth/domain/account"
	emaildomain "nfxidentity/modules/auth/domain/email"
	identitydomain "nfxidentity/modules/auth/domain/identity"
	refreshtokendomain "nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/pkgs/ptrx"
	"nfxidentity/pkgs/tokenx"
	"nfxidentity/pkgs/tokenx/hashx"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ProfileItem struct {
	ProfileID     uuid.UUID              `json:"profile_id"`
	Kind          enums.AuthProfileScope `json:"kind"`
	Roles         []string               `json:"roles"`
	DisplayName   *string                `json:"display_name"`
	AvatarImageID *uuid.UUID             `json:"avatar_image_id"`
	City          *string                `json:"city"`
	Country       *string                `json:"country"`
}

type WithEmailInput struct {
	Email    string
	Password string
	DeviceID string
}

type WithEmailOutput struct {
	AccountID    uuid.UUID     `json:"account_id"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	Profiles     []ProfileItem `json:"profiles"`
}

func (s *Service) WithEmail(ctx context.Context, in WithEmailInput) (*WithEmailOutput, error) {
	var account *accountdomain.Account
	var email *emaildomain.Email
	var identity *identitydomain.Identity
	var refreshTokenEntity *refreshtokendomain.RefreshToken
	var accessToken string
	var refreshToken string
	var refreshClaims *tokenx.TokenClaims
	loginAt := time.Now().UTC()

	if err := s.txManager.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		emailRepo := s.repoFactory.Email(uow)
		accountRepo := s.repoFactory.Account(uow)
		identityRepo := s.repoFactory.Identity(uow)
		refreshTokenRepo := s.repoFactory.RefreshToken(uow)
		var err error

		if email, err = emailRepo.Get.ByEmail(ctx, in.Email); err != nil {
			return authErr.ErrEmailOrPasswordIncorrect
		}
		if !email.IsVerified() {
			return authErr.ErrEmailOrPasswordIncorrect
		}

		if account, err = accountRepo.Get.ByID(ctx, email.AccountID()); err != nil {
			return authErr.ErrEmailOrPasswordIncorrect
		}
		if !account.IsActive() {
			return authErr.ErrAccountInactive
		}

		identities, err := identityRepo.Get.ByAccountID(ctx, account.ID())
		if err != nil {
			return authErr.ErrEmailOrPasswordIncorrect
		}
		for _, id := range identities {
			if id.IdentityProvider() == enums.AuthIdentityProviderPassword {
				identity = id
				break
			}
		}
		if identity == nil || identity.PasswordHash() == nil {
			return authErr.ErrEmailOrPasswordIncorrect
		}
		if err := bcrypt.CompareHashAndPassword([]byte(*identity.PasswordHash()), []byte(in.Password)); err != nil {
			return authErr.ErrEmailOrPasswordIncorrect
		}

		select {
		case <-time.After(time.Second):
		case <-ctx.Done():
			return ctx.Err()
		}

		allEmails, err := emailRepo.Get.ByAccountID(ctx, account.ID())
		if err != nil {
			return authErr.ErrEmailOrPasswordIncorrect
		}
		for _, e := range allEmails {
			if !e.IsPrimary() {
				continue
			}
			if identity.ProviderSubject() == e.Email() {
				break
			}
			if err := identity.UpdateProviderSubject(e.Email()); err != nil {
				return authErr.ErrLoginFailed.WithCause(err)
			}
			if err := identityRepo.Update.Generic(ctx, identity); err != nil {
				return authErr.ErrLoginFailed.WithCause(err)
			}
			break
		}

		if accessToken, refreshToken, err = s.tokenx.GenerateTokenPair(
			account.ID().String(),
			"",
			in.Email,
			in.Email,
			"",
			"",
		); err != nil {
			return authErr.ErrLoginFailed.WithCause(err)
		}

		if refreshClaims, err = s.tokenx.VerifyRefreshToken(refreshToken); err != nil || refreshClaims == nil || refreshClaims.ExpiresAt == nil {
			return authErr.ErrLoginFailed.WithCause(err)
		}

		if refreshTokenEntity, err = refreshtokendomain.NewRefreshToken(refreshtokendomain.NewRefreshTokenParams{
			AccountID:  account.ID(),
			IdentityID: ptrx.Ptr(identity.ID()),
			DeviceID:   ptrx.Ptr(in.DeviceID),
			TokenHash:  hashx.SHA256HexString(refreshToken),
			ExpiresAt:  refreshClaims.ExpiresAt.Time,
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
			loginAt,
		); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	profiles := make([]ProfileItem, 0)

	communityList, err := s.profileQuery.ForgerList.ByAccountID(ctx, account.ID())
	if err != nil {
		return nil, authErr.ErrForgerProfileNotFound.WithCause(err)
	}
	for _, p := range communityList {
		profiles = append(profiles, ProfileItem{
			ProfileID:     p.ProfileID,
			Kind:          enums.AuthProfileScopeCommunity,
			Roles:         forgerRolesToStrings(p.ForgerRoles),
			DisplayName:   p.DisplayName,
			AvatarImageID: p.AvatarImageID,
			City:          p.City,
			Country:       p.Country,
		})
	}

	authorityList, err := s.profileQuery.AuthorityList.ByAccountID(ctx, account.ID())
	if err != nil {
		return nil, authErr.ErrAuthorityProfileNotFound.WithCause(err)
	}
	for _, p := range authorityList {
		profiles = append(profiles, ProfileItem{
			ProfileID:     p.ProfileID,
			Kind:          enums.AuthProfileScopeAuthority,
			Roles:         authorityRolesToStrings(p.AuthorityRoles),
			DisplayName:   p.DisplayName,
			AvatarImageID: p.AvatarImageID,
			City:          p.City,
			Country:       p.Country,
		})
	}

	if len(profiles) == 0 {
		return nil, authErr.ErrForgerProfileNotFound
	}

	return &WithEmailOutput{
		AccountID:    account.ID(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Profiles:     profiles,
	}, nil
}

func forgerRolesToStrings(roles []enums.AuthForgerRole) []string {
	out := make([]string, 0, len(roles))
	for _, role := range roles {
		out = append(out, string(role))
	}
	return out
}

func authorityRolesToStrings(roles []enums.AuthAuthorityRole) []string {
	out := make([]string, 0, len(roles))
	for _, role := range roles {
		out = append(out, string(role))
	}
	return out
}
