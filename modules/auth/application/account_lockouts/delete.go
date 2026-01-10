package account_lockouts

import (
	"context"
	accountLockoutCommands "nfxid/modules/auth/application/account_lockouts/commands"
)

// DeleteAccountLockout 删除账户锁定
func (s *Service) DeleteAccountLockout(ctx context.Context, cmd accountLockoutCommands.DeleteAccountLockoutCmd) error {
	// Delete from repository (hard delete)
	return s.accountLockoutRepo.Delete.ByUserIDAndTenantID(ctx, cmd.UserID, cmd.TenantID)
}
