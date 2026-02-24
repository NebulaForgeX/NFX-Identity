package auth

import "nfxid/pkgs/errx"

const (
	CodeAccountLockoutNotFound = "ACCOUNT_LOCKOUT_NOT_FOUND"
	CodeLockReasonRequired     = "LOCK_REASON_REQUIRED"
	CodeInvalidLockReason      = "INVALID_LOCK_REASON"
	CodeAccountAlreadyLocked   = "ACCOUNT_ALREADY_LOCKED"
	CodeAccountNotLocked       = "ACCOUNT_NOT_LOCKED"
)

var (
	ErrAccountLockoutNotFound = errx.NotFound(CodeAccountLockoutNotFound, "account lockout not found")
	ErrLockReasonRequired     = errx.InvalidArg(CodeLockReasonRequired, "lock reason is required")
	ErrInvalidLockReason      = errx.InvalidArg(CodeInvalidLockReason, "invalid lock reason")
	ErrAccountAlreadyLocked   = errx.Conflict(CodeAccountAlreadyLocked, "account already locked")
	ErrAccountNotLocked       = errx.FailedPrecond(CodeAccountNotLocked, "account is not locked")
)
