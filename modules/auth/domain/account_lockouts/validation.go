package account_lockouts

import (
	authErr "nfxid/errors/src/auth"

	"github.com/google/uuid"
)

func (al *AccountLockout) Validate() error {
	if al.UserID() == uuid.Nil {
		return authErr.ErrUserIDRequired
	}
	if al.LockReason() == "" {
		return authErr.ErrLockReasonRequired
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
		return authErr.ErrInvalidLockReason
	}
	return nil
}
