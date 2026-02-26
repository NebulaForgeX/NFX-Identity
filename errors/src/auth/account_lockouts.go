package auth

import "nfxidentity/pkgs/errx"

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

/*
!ACCOUNT_LOCKOUT_NOT_FOUND
*en<account lockout not found>
*zh<账户锁定记录不存在>
*fr<verrouillage de compte introuvable>

!LOCK_REASON_REQUIRED
*en<lock reason required>
*zh<锁定原因为必填>
*fr<raison du verrouillage requise>

!INVALID_LOCK_REASON
*en<invalid lock reason>
*zh<无效的锁定原因>
*fr<raison de verrouillage invalide>

!ACCOUNT_ALREADY_LOCKED
*en<account already locked>
*zh<账户已锁定>
*fr<compte déjà verrouillé>

!ACCOUNT_NOT_LOCKED
*en<account not locked>
*zh<账户未锁定>
*fr<compte non verrouillé>

*/
