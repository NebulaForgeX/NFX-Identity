package account_lockouts

import "nfxid/pkgs/errx"

var (
	ErrAccountLockoutNotFound = errx.NotFound("ACCOUNT_LOCKOUT_NOT_FOUND", "account lockout not found")
	ErrUserIDRequired         = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrLockReasonRequired     = errx.InvalidArg("LOCK_REASON_REQUIRED", "lock reason is required")
	ErrInvalidLockReason      = errx.InvalidArg("INVALID_LOCK_REASON", "invalid lock reason")
	ErrAccountAlreadyLocked   = errx.Conflict("ACCOUNT_ALREADY_LOCKED", "account already locked")
	ErrAccountNotLocked       = errx.FailedPrecond("ACCOUNT_NOT_LOCKED", "account is not locked")
)
