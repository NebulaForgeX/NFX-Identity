package refresh_tokens

import (
	"context"
	refreshTokenCommands "nfxid/modules/auth/application/refresh_tokens/commands"
)

// DeleteRefreshToken 删除刷新令牌
func (s *Service) DeleteRefreshToken(ctx context.Context, cmd refreshTokenCommands.DeleteRefreshTokenCmd) error {
	// Delete from repository (hard delete)
	return s.refreshTokenRepo.Delete.ByID(ctx, cmd.RefreshTokenID)
}
