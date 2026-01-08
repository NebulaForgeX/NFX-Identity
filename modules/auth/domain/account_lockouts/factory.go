package account_lockouts

import (
	"time"

	"github.com/google/uuid"
)

type NewAccountLockoutParams struct {
	UserID      uuid.UUID
	TenantID    uuid.UUID
	LockedUntil *time.Time
	LockReason  LockReason
	LockedBy    *string
	ActorID     *uuid.UUID
}

func NewAccountLockout(p NewAccountLockoutParams) (*AccountLockout, error) {
	if err := validateAccountLockoutParams(p); err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewAccountLockoutFromState(AccountLockoutState{
		UserID:      p.UserID,
		TenantID:    p.TenantID,
		LockedUntil: p.LockedUntil,
		LockReason:  p.LockReason,
		LockedAt:    now,
		LockedBy:    p.LockedBy,
		ActorID:     p.ActorID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}), nil
}

func NewAccountLockoutFromState(st AccountLockoutState) *AccountLockout {
	return &AccountLockout{state: st}
}

func validateAccountLockoutParams(p NewAccountLockoutParams) error {
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if p.TenantID == uuid.Nil {
		return ErrTenantIDRequired
	}
	if p.LockReason == "" {
		return ErrLockReasonRequired
	}
	validReasons := map[LockReason]struct{}{
		LockReasonTooManyAttempts:    {},
		LockReasonAdminLock:          {},
		LockReasonRiskDetected:       {},
		LockReasonSuspiciousActivity: {},
		LockReasonCompliance:         {},
		LockReasonOther:              {},
	}
	if _, ok := validReasons[p.LockReason]; !ok {
		return ErrInvalidLockReason
	}
	return nil
}
