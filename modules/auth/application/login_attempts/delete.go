package login_attempts

import (
	"context"
	loginAttemptCommands "nfxid/modules/auth/application/login_attempts/commands"
)

// DeleteLoginAttempt 删除登录尝试（如果需要清理历史记录）
func (s *Service) DeleteLoginAttempt(ctx context.Context, cmd loginAttemptCommands.DeleteLoginAttemptCmd) error {
	// Delete from repository (hard delete)
	return s.loginAttemptRepo.Delete.ByID(ctx, cmd.LoginAttemptID)
}
