package auth

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInvalidRefreshToken = errors.New("invalid or expired refresh token")
	ErrAccountLocked       = errors.New("account is locked due to too many failed login attempts")
	ErrEmailAlreadyExists  = errors.New("email already exists")
	ErrEmailAlreadyVerified = errors.New("email already verified, please login")
	ErrInvalidVerificationCode = errors.New("invalid or expired verification code")
)
