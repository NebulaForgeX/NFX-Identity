package account_lockouts

import "github.com/google/uuid"

func (al *AccountLockout) Validate() error {
	if al.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	if al.LockReason() == "" {
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
	if _, ok := validReasons[al.LockReason()]; !ok {
		return ErrInvalidLockReason
	}
	return nil
}
