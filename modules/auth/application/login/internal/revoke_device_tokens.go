package internal

import (
	"context"
	"errors"
	"time"

	authErr "nfxidentity/errors/src/auth"
	refreshtokendomain "nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/pkgs/logx"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// RevokeOtherRefreshTokensForDevice revokes every active refresh token for
// account+device except keepID. Runs in the same UoW as minting the new token
// (Auth0-style: one current refresh per device; no async race with Kafka).
func RevokeOtherRefreshTokensForDevice(
	ctx context.Context,
	repo *refreshtokendomain.Repo,
	accountID uuid.UUID,
	deviceID string,
	keepID uuid.UUID,
	at time.Time,
) error {
	existing, err := repo.Get.ByAccountIDAndDeviceID(ctx, accountID, deviceID)
	if err != nil {
		return authErr.ErrLoginFailed.WithCause(err)
	}
	for _, token := range existing {
		if token == nil || token.ID() == keepID {
			continue
		}
		if err := token.Revoke(at); err != nil {
			if errors.Is(err, authErr.ErrRefreshTokenAlreadyRevoked) {
				continue
			}
			logx.L().Warn("revoke other refresh tokens: revoke failed",
				zap.String("account_id", accountID.String()),
				zap.String("device_id", deviceID),
				zap.String("refresh_token_id", token.ID().String()),
				zap.Error(err),
			)
			continue
		}
		if err := repo.Update.Generic(ctx, token); err != nil {
			return authErr.ErrLoginFailed.WithCause(err)
		}
	}
	return nil
}
