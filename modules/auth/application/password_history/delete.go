package password_history

import (
	"context"
	passwordHistoryCommands "nfxid/modules/auth/application/password_history/commands"
)

// DeletePasswordHistory 删除密码历史
func (s *Service) DeletePasswordHistory(ctx context.Context, cmd passwordHistoryCommands.DeletePasswordHistoryCmd) error {
	// Delete from repository (hard delete)
	return s.passwordHistoryRepo.Delete.ByID(ctx, cmd.PasswordHistoryID)
}
