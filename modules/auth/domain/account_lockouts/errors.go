package account_lockouts

import "errors"

var (
	ErrAccountLockoutNotFound   = errors.New("account lockout not found")
	ErrUserIDRequired           = errors.New("user id is required")
	ErrTenantIDRequired         = errors.New("tenant id is required")
	ErrLockReasonRequired       = errors.New("lock reason is required")
	ErrInvalidLockReason        = errors.New("invalid lock reason")
	ErrAccountAlreadyLocked     = errors.New("account already locked")
	ErrAccountNotLocked         = errors.New("account is not locked")
)
