package refresh_tokens

import (
	"context"
	refreshTokenCommands "nfxid/modules/auth/application/refresh_tokens/commands"
)

// RevokeRefreshToken 撤销刷新令牌
func (s *Service) RevokeRefreshToken(ctx context.Context, cmd refreshTokenCommands.RevokeRefreshTokenCmd) error {
	// Get domain entity
	refreshToken, err := s.refreshTokenRepo.Get.ByTokenID(ctx, cmd.TokenID)
	if err != nil {
		return err
	}

	// Revoke domain entity
	if err := refreshToken.Revoke(cmd.RevokeReason); err != nil {
		return err
	}

	// Save to repository
	return s.refreshTokenRepo.Update.Revoke(ctx, cmd.TokenID, cmd.RevokeReason)
}
