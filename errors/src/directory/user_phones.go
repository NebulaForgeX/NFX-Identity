package directory

import "nfxid/pkgs/errx"

const (
	CodeUserPhoneNotFound       = "USER_PHONE_NOT_FOUND"
	CodePhoneRequired           = "PHONE_REQUIRED"
	CodePhoneAlreadyExists      = "PHONE_ALREADY_EXISTS"
	CodeInvalidPhone            = "INVALID_PHONE"
	CodeVerificationCodeExpired = "VERIFICATION_CODE_EXPIRED"
)

var (
	ErrUserPhoneNotFound       = errx.NotFound(CodeUserPhoneNotFound, "user phone not found")
	ErrPhoneRequired           = errx.InvalidArg(CodePhoneRequired, "phone is required")
	ErrPhoneAlreadyExists      = errx.Conflict(CodePhoneAlreadyExists, "phone already exists")
	ErrInvalidPhone            = errx.InvalidArg(CodeInvalidPhone, "invalid phone format")
	ErrVerificationCodeExpired = errx.Expired(CodeVerificationCodeExpired, "verification code expired")
)
