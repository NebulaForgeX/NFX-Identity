package auth

import "nfxid/pkgs/errx"

const (
	CodePasswordResetNotFound = "PASSWORD_RESET_NOT_FOUND"
	CodeResetIDRequired       = "RESET_ID_REQUIRED"
	CodeDeliveryRequired      = "DELIVERY_REQUIRED"
	CodeCodeHashRequired      = "CODE_HASH_REQUIRED"
	CodeResetIDAlreadyExists  = "RESET_ID_ALREADY_EXISTS"
	CodeInvalidResetDelivery  = "INVALID_RESET_DELIVERY"
	CodeInvalidResetStatus    = "INVALID_RESET_STATUS"
	CodeResetAlreadyUsed      = "RESET_ALREADY_USED"
	CodeResetExpired          = "RESET_EXPIRED"
)

var (
	ErrPasswordResetNotFound = errx.NotFound(CodePasswordResetNotFound, "password reset not found")
	ErrResetIDRequired       = errx.InvalidArg(CodeResetIDRequired, "reset id is required")
	ErrDeliveryRequired      = errx.InvalidArg(CodeDeliveryRequired, "delivery is required")
	ErrCodeHashRequired      = errx.InvalidArg(CodeCodeHashRequired, "code hash is required")
	ErrResetIDAlreadyExists  = errx.Conflict(CodeResetIDAlreadyExists, "reset id already exists")
	ErrInvalidResetDelivery  = errx.InvalidArg(CodeInvalidResetDelivery, "invalid reset delivery")
	ErrInvalidResetStatus    = errx.InvalidArg(CodeInvalidResetStatus, "invalid reset status")
	ErrResetAlreadyUsed      = errx.Conflict(CodeResetAlreadyUsed, "reset already used")
	ErrResetExpired          = errx.Expired(CodeResetExpired, "reset expired")
)
