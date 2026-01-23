package account_lockouts

import (
	"time"

	"github.com/google/uuid"
)

type LockReason string

const (
	LockReasonTooManyAttempts    LockReason = "too_many_attempts"
	LockReasonAdminLock          LockReason = "admin_lock"
	LockReasonRiskDetected       LockReason = "risk_detected"
	LockReasonSuspiciousActivity LockReason = "suspicious_activity"
	LockReasonCompliance         LockReason = "compliance"
	LockReasonOther              LockReason = "other"
)

type AccountLockout struct {
	state AccountLockoutState
}

type AccountLockoutState struct {
	UserID        uuid.UUID
	LockedUntil   *time.Time
	LockReason    LockReason
	LockedAt      time.Time
	LockedBy      *string
	ActorID       *uuid.UUID
	UnlockedAt    *time.Time
	UnlockedBy    *string
	UnlockActorID *uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (al *AccountLockout) UserID() uuid.UUID      { return al.state.UserID }
func (al *AccountLockout) LockedUntil() *time.Time { return al.state.LockedUntil }
func (al *AccountLockout) LockReason() LockReason { return al.state.LockReason }
func (al *AccountLockout) LockedAt() time.Time    { return al.state.LockedAt }
func (al *AccountLockout) LockedBy() *string      { return al.state.LockedBy }
func (al *AccountLockout) ActorID() *uuid.UUID    { return al.state.ActorID }
func (al *AccountLockout) UnlockedAt() *time.Time { return al.state.UnlockedAt }
func (al *AccountLockout) UnlockedBy() *string    { return al.state.UnlockedBy }
func (al *AccountLockout) UnlockActorID() *uuid.UUID { return al.state.UnlockActorID }
func (al *AccountLockout) CreatedAt() time.Time   { return al.state.CreatedAt }
func (al *AccountLockout) UpdatedAt() time.Time   { return al.state.UpdatedAt }
