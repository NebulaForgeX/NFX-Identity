package user_phones

import "nfxid/pkgs/errx"

var (
	ErrUserPhoneNotFound        = errx.NotFound("USER_PHONE_NOT_FOUND", "user phone not found")
	ErrUserIDRequired           = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrPhoneRequired            = errx.InvalidArg("PHONE_REQUIRED", "phone is required")
	ErrPhoneAlreadyExists       = errx.Conflict("PHONE_ALREADY_EXISTS", "phone already exists")
	ErrInvalidPhone             = errx.InvalidArg("INVALID_PHONE", "invalid phone format")
	ErrVerificationCodeExpired  = errx.Expired("VERIFICATION_CODE_EXPIRED", "verification code expired")
)
