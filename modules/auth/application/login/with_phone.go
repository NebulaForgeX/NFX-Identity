package login

import (
	"context"
	"strings"
	"time"

	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	logininternal "nfxidentity/modules/auth/application/login/internal"
	accountdomain "nfxidentity/modules/auth/domain/account"
	identitydomain "nfxidentity/modules/auth/domain/identity"
	refreshtokendomain "nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/pkgs/ptrx"
	"nfxidentity/pkgs/tokenx"
	"nfxidentity/pkgs/tokenx/hashx"
	"nfxidentity/pkgs/transaction"

	"golang.org/x/crypto/bcrypt"
)

type WithPhoneInput struct {
	Phone    string
	Password string
	DeviceID string
}

func (s *Service) WithPhone(ctx context.Context, in WithPhoneInput) (*WithEmailOutput, error) {
	var account *accountdomain.Account
	var identity *identitydomain.Identity
	var refreshTokenEntity *refreshtokendomain.RefreshToken
	var accessToken string
	var refreshToken string
	var refreshClaims *tokenx.TokenClaims
	loginAt := time.Now().UTC()
	phoneNum := strings.TrimSpace(in.Phone)

	if err := s.txManager.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		phoneRepo := s.repoFactory.Phone(uow)
		accountRepo := s.repoFactory.Account(uow)
		identityRepo := s.repoFactory.Identity(uow)
		refreshTokenRepo := s.repoFactory.RefreshToken(uow)
		var err error

		row, err := phoneRepo.Get.ByPhone(ctx, phoneNum)
		if err != nil {
			return authErr.ErrInvalidCredentials
		}
		if account, err = accountRepo.Get.ByID(ctx, row.AccountID()); err != nil {
			return authErr.ErrInvalidCredentials
		}
		if !account.IsActive() {
			return authErr.ErrAccountInactive
		}

		identities, err := identityRepo.Get.ByAccountID(ctx, account.ID())
		if err != nil {
			return authErr.ErrInvalidCredentials
		}
		for _, id := range identities {
			if id.IdentityProvider() == enums.AuthIdentityProviderPassword {
				identity = id
				break
			}
		}
		if identity == nil || identity.PasswordHash() == nil {
			return authErr.ErrInvalidCredentials
		}
		if err := bcrypt.CompareHashAndPassword([]byte(*identity.PasswordHash()), []byte(in.Password)); err != nil {
			return authErr.ErrInvalidCredentials
		}

		if accessToken, refreshToken, err = s.tokenx.GenerateTokenPair(
			account.ID().String(),
			"",
			phoneNum,
			"",
			phoneNum,
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
		return logininternal.RevokeOtherRefreshTokensForDevice(
			ctx,
			refreshTokenRepo,
			account.ID(),
			in.DeviceID,
			refreshTokenEntity.ID(),
			loginAt,
		)
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
