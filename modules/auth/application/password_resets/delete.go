package password_resets

import (
	"context"
	passwordResetCommands "nfxid/modules/auth/application/password_resets/commands"
)

// DeletePasswordReset 删除密码重置
func (s *Service) DeletePasswordReset(ctx context.Context, cmd passwordResetCommands.DeletePasswordResetCmd) error {
	// Delete from repository (hard delete)
	return s.passwordResetRepo.Delete.ByResetID(ctx, cmd.ResetID)
}
