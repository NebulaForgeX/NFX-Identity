package account_lockouts

import (
	"context"
	accountLockoutCommands "nfxid/modules/auth/application/account_lockouts/commands"
)

// UnlockAccount 解锁账户
func (s *Service) UnlockAccount(ctx context.Context, cmd accountLockoutCommands.UnlockAccountCmd) error {
	// Get domain entity
	accountLockout, err := s.accountLockoutRepo.Get.ByUserIDAndTenantID(ctx, cmd.UserID, cmd.TenantID)
	if err != nil {
		return err
	}

	// Unlock domain entity
	if err := accountLockout.Unlock(cmd.UnlockedBy, cmd.UnlockActorID); err != nil {
		return err
	}

	// Save to repository
	return s.accountLockoutRepo.Update.Generic(ctx, accountLockout)
}
