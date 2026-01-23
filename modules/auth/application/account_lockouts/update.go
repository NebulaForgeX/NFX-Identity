package account_lockouts

import (
	"context"
	accountLockoutCommands "nfxid/modules/auth/application/account_lockouts/commands"
)

// UnlockAccount 解锁账户
func (s *Service) UnlockAccount(ctx context.Context, cmd accountLockoutCommands.UnlockAccountCmd) error {
	// Unlock via repository
	return s.accountLockoutRepo.Update.Unlock(ctx, cmd.UserID, cmd.UnlockedBy, cmd.UnlockActorID)
}
