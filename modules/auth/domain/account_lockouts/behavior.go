package account_lockouts

import (
	"time"

	"github.com/google/uuid"
)

func (al *AccountLockout) Unlock(unlockedBy string, unlockActorID *uuid.UUID) error {
	if al.UnlockedAt() != nil {
		return nil // already unlocked, idempotent
	}

	now := time.Now().UTC()
	al.state.UnlockedAt = &now
	al.state.UnlockedBy = &unlockedBy
	al.state.UnlockActorID = unlockActorID
	al.state.LockedUntil = nil
	al.state.UpdatedAt = now
	return nil
}

func (al *AccountLockout) IsLocked() bool {
	if al.UnlockedAt() != nil {
		return false
	}
	if al.LockedUntil() == nil {
		return true // permanently locked
	}
	return time.Now().UTC().Before(*al.LockedUntil())
}

func (al *AccountLockout) UpdateLockedUntil(lockedUntil *time.Time) error {
	al.state.LockedUntil = lockedUntil
	al.state.UpdatedAt = time.Now().UTC()
	return nil
}
