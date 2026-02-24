package directory

import "nfxid/pkgs/errx"

const (
	CodeUserEmailNotFound  = "USER_EMAIL_NOT_FOUND"
	CodeEmailRequired      = "EMAIL_REQUIRED"
	CodeEmailAlreadyExists = "EMAIL_ALREADY_EXISTS"
	CodeInvalidEmail       = "INVALID_EMAIL"
)

var (
	ErrUserEmailNotFound  = errx.NotFound(CodeUserEmailNotFound, "user email not found")
	ErrEmailRequired      = errx.InvalidArg(CodeEmailRequired, "email is required")
	ErrEmailAlreadyExists = errx.Conflict(CodeEmailAlreadyExists, "email already exists")
	ErrInvalidEmail       = errx.InvalidArg(CodeInvalidEmail, "invalid email format")
)
