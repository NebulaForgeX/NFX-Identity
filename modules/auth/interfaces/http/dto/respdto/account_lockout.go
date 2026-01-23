package respdto

import (
	"time"

	accountLockoutAppResult "nfxid/modules/auth/application/account_lockouts/results"

	"github.com/google/uuid"
)

type AccountLockoutDTO struct {
	UserID        uuid.UUID  `json:"user_id"`
	LockedUntil   *time.Time `json:"locked_until,omitempty"`
	LockReason    string     `json:"lock_reason"`
	LockedAt      time.Time  `json:"locked_at"`
	LockedBy      *string    `json:"locked_by,omitempty"`
	ActorID       *uuid.UUID `json:"actor_id,omitempty"`
	UnlockedAt    *time.Time `json:"unlocked_at,omitempty"`
	UnlockedBy    *string    `json:"unlocked_by,omitempty"`
	UnlockActorID *uuid.UUID `json:"unlock_actor_id,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// AccountLockoutROToDTO converts application AccountLockoutRO to response DTO
func AccountLockoutROToDTO(v *accountLockoutAppResult.AccountLockoutRO) *AccountLockoutDTO {
	if v == nil {
		return nil
	}

	return &AccountLockoutDTO{
		UserID:        v.UserID,
		LockedUntil:   v.LockedUntil,
		LockReason:    string(v.LockReason),
		LockedAt:      v.LockedAt,
		LockedBy:      v.LockedBy,
		ActorID:       v.ActorID,
		UnlockedAt:    v.UnlockedAt,
		UnlockedBy:    v.UnlockedBy,
		UnlockActorID: v.UnlockActorID,
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     v.UpdatedAt,
	}
}
