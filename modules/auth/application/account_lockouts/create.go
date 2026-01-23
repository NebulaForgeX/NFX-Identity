package account_lockouts

import (
	"context"
	"time"
	accountLockoutCommands "nfxid/modules/auth/application/account_lockouts/commands"
	accountLockoutDomain "nfxid/modules/auth/domain/account_lockouts"
)

// CreateAccountLockout 创建账户锁定
func (s *Service) CreateAccountLockout(ctx context.Context, cmd accountLockoutCommands.CreateAccountLockoutCmd) error {
	// Check if account is already locked
	if exists, _ := s.accountLockoutRepo.Check.ByUserID(ctx, cmd.UserID); exists {
		return accountLockoutDomain.ErrAccountAlreadyLocked
	}

	var lockedUntil *time.Time
	if cmd.LockedUntil != nil && *cmd.LockedUntil != "" {
		parsed, err := time.Parse(time.RFC3339, *cmd.LockedUntil)
		if err != nil {
			return err
		}
		lockedUntil = &parsed
	}

	// Create domain entity
	accountLockout, err := accountLockoutDomain.NewAccountLockout(accountLockoutDomain.NewAccountLockoutParams{
		UserID:      cmd.UserID,
		LockedUntil: lockedUntil,
		LockReason:  cmd.LockReason,
		LockedBy:    cmd.LockedBy,
		ActorID:     cmd.ActorID,
	})
	if err != nil {
		return err
	}

	// Save to repository
	return s.accountLockoutRepo.Create.New(ctx, accountLockout)
}
