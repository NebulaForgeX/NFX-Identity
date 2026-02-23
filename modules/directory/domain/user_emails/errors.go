package user_emails

import "nfxid/pkgs/errx"

var (
	ErrUserEmailNotFound  = errx.NotFound("USER_EMAIL_NOT_FOUND", "user email not found")
	ErrUserIDRequired     = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrEmailRequired      = errx.InvalidArg("EMAIL_REQUIRED", "email is required")
	ErrEmailAlreadyExists = errx.Conflict("EMAIL_ALREADY_EXISTS", "email already exists")
	ErrInvalidEmail       = errx.InvalidArg("INVALID_EMAIL", "invalid email format")
)
