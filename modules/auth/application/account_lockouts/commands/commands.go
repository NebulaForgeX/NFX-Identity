package commands

import (
	"nfxid/modules/auth/domain/account_lockouts"

	"github.com/google/uuid"
)

// CreateAccountLockoutCmd 创建账户锁定命令
type CreateAccountLockoutCmd struct {
	UserID      uuid.UUID
	TenantID    uuid.UUID
	LockedUntil *string
	LockReason  account_lockouts.LockReason
	LockedBy    *string
	ActorID     *uuid.UUID
}

// UnlockAccountCmd 解锁账户命令
type UnlockAccountCmd struct {
	UserID        uuid.UUID
	TenantID      uuid.UUID
	UnlockedBy    string
	UnlockActorID *uuid.UUID
}

// DeleteAccountLockoutCmd 删除账户锁定命令
type DeleteAccountLockoutCmd struct {
	UserID   uuid.UUID
	TenantID uuid.UUID
}
