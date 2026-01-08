package user_phones

import "errors"

var (
	ErrUserPhoneNotFound   = errors.New("user phone not found")
	ErrUserIDRequired      = errors.New("user id is required")
	ErrPhoneRequired       = errors.New("phone is required")
	ErrPhoneAlreadyExists  = errors.New("phone already exists")
	ErrInvalidPhone        = errors.New("invalid phone format")
	ErrVerificationCodeExpired = errors.New("verification code expired")
)
