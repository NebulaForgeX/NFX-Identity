package login

import (
	"context"
	"strings"
	"time"

	"nfxidentity/constants"
	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	sysErr "nfxidentity/errors/src/sys"
	logininternal "nfxidentity/modules/auth/application/login/internal"
	refreshtokendomain "nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/pkgs/logx"
	"nfxidentity/pkgs/tokenx"
	"nfxidentity/pkgs/tokenx/hashx"
	"nfxidentity/pkgs/transaction"

	"go.uber.org/zap"
)

type RefreshInput struct {
	RefreshToken string
	DeviceID     string
}

type RefreshOutput struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *Service) Refresh(ctx context.Context, in RefreshInput) (*RefreshOutput, error) {
	refreshClaims, err := s.tokenx.VerifyRefreshToken(in.RefreshToken)
	if err != nil {
		return nil, authErr.ErrInvalidRefreshToken.WithCause(err)
	}

	tokenHash := hashx.SHA256HexString(in.RefreshToken)
	now := time.Now().UTC()

	var accessToken, newRefreshToken string
	var newRefreshClaims *tokenx.TokenClaims

	if err := s.txManager.WithUoW(ctx, func(ctx context.Context, uow transaction.UoW) error {
		refreshTokenRepo := s.repoFactory.RefreshToken(uow)

		stored, err := refreshTokenRepo.Get.ByTokenHash(ctx, tokenHash)
		if err != nil {
			return authErr.ErrInvalidRefreshToken.WithCause(err)
		}
		if stored.IsRevoked() || stored.IsExpired(now) {
			return authErr.ErrInvalidRefreshToken
		}
		if stored.AccountID().String() != refreshClaims.AccountID {
			return authErr.ErrInvalidRefreshToken
		}

		deviceID, err := logininternal.ResolveRefreshDeviceID(stored.DeviceID(), in.DeviceID)
		if err != nil {
			return err
		}

		if err := stored.Revoke(now); err != nil {
			return err
		}
		if err := refreshTokenRepo.Update.Generic(ctx, stored); err != nil {
			return err
		}
		profileScope := strings.TrimSpace(refreshClaims.ProfileScope)
		if profileScope != "" && !constants.AuthProfileScope.Valid(enums.AuthProfileScope(profileScope)) {
			return sysErr.ErrTokenInvalidProfileScope.WithDetail("profileScope", profileScope)
		}

		accessToken, newRefreshToken, err = s.tokenx.GenerateTokenPair(
			refreshClaims.AccountID,
			refreshClaims.ProfileID,
			refreshClaims.Username,
			refreshClaims.Email,
			refreshClaims.Phone,
			profileScope,
		)
		if err != nil {
			return authErr.ErrInvalidRefreshToken.WithCause(err)
		}

		newRefreshClaims, err = s.tokenx.VerifyRefreshToken(newRefreshToken)
		if err != nil || newRefreshClaims == nil || newRefreshClaims.ExpiresAt == nil {
			return authErr.ErrInvalidRefreshToken.WithCause(err)
		}

		newEntity, err := refreshtokendomain.NewRefreshToken(refreshtokendomain.NewRefreshTokenParams{
			AccountID:  stored.AccountID(),
			IdentityID: stored.IdentityID(),
			DeviceID:   deviceID,
			TokenHash:  hashx.SHA256HexString(newRefreshToken),
			ExpiresAt:  newRefreshClaims.ExpiresAt.Time,
		})
		if err != nil {
			return err
		}
		return refreshTokenRepo.Create.New(ctx, newEntity)
	}); err != nil {
		return nil, err
	}

	logx.L().Info("auth token refreshed",
		zap.String("account_id", refreshClaims.AccountID),
	)

	return &RefreshOutput{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
