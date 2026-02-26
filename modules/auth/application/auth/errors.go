package auth

import "nfxidentity/pkgs/errx"

var (
	ErrInvalidCredentials      = errx.Unauthorized("INVALID_CREDENTIALS", "invalid email or password")
	ErrInvalidRefreshToken     = errx.Unauthorized("INVALID_REFRESH_TOKEN", "invalid or expired refresh token")
	ErrAccountLocked           = errx.Forbidden("ACCOUNT_LOCKED", "account is locked due to too many failed login attempts")
	ErrEmailAlreadyExists      = errx.Conflict("EMAIL_ALREADY_EXISTS", "email already exists")
	ErrEmailAlreadyVerified    = errx.Conflict("EMAIL_ALREADY_VERIFIED", "email already verified, please login")
	ErrInvalidVerificationCode = errx.InvalidArg("INVALID_VERIFICATION_CODE", "invalid or expired verification code")
)
