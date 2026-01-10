package results

import (
	"time"

	"nfxid/modules/auth/domain/account_lockouts"

	"github.com/google/uuid"
)

type AccountLockoutRO struct {
	UserID        uuid.UUID
	TenantID      uuid.UUID
	LockedUntil   *time.Time
	LockReason    account_lockouts.LockReason
	LockedAt      time.Time
	LockedBy      *string
	ActorID       *uuid.UUID
	UnlockedAt    *time.Time
	UnlockedBy    *string
	UnlockActorID *uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// AccountLockoutMapper 将 Domain AccountLockout 转换为 Application AccountLockoutRO
func AccountLockoutMapper(al *account_lockouts.AccountLockout) AccountLockoutRO {
	if al == nil {
		return AccountLockoutRO{}
	}

	return AccountLockoutRO{
		UserID:        al.UserID(),
		TenantID:      al.TenantID(),
		LockedUntil:   al.LockedUntil(),
		LockReason:    al.LockReason(),
		LockedAt:      al.LockedAt(),
		LockedBy:      al.LockedBy(),
		ActorID:       al.ActorID(),
		UnlockedAt:    al.UnlockedAt(),
		UnlockedBy:    al.UnlockedBy(),
		UnlockActorID: al.UnlockActorID(),
		CreatedAt:     al.CreatedAt(),
		UpdatedAt:     al.UpdatedAt(),
	}
}
