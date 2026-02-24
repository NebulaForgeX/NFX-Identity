package directory

import "nfxid/pkgs/errx"

const (
	CodeUserNotFound          = "USER_NOT_FOUND"
	CodeUsernameRequired      = "USERNAME_REQUIRED"
	CodeUsernameAlreadyExists = "USERNAME_ALREADY_EXISTS"
	CodeInvalidUserStatus     = "INVALID_USER_STATUS"
)

var (
	ErrUserNotFound          = errx.NotFound(CodeUserNotFound, "user not found")
	ErrUsernameRequired      = errx.InvalidArg(CodeUsernameRequired, "username is required")
	ErrUsernameAlreadyExists = errx.Conflict(CodeUsernameAlreadyExists, "username already exists")
	ErrInvalidUserStatus     = errx.InvalidArg(CodeInvalidUserStatus, "invalid user status")
)
