package users

import "nfxid/pkgs/errx"

var (
	ErrUserNotFound         = errx.NotFound("USER_NOT_FOUND", "user not found")
	ErrUsernameRequired     = errx.InvalidArg("USERNAME_REQUIRED", "username is required")
	ErrUsernameAlreadyExists = errx.Conflict("USERNAME_ALREADY_EXISTS", "username already exists")
	ErrInvalidUserStatus    = errx.InvalidArg("INVALID_USER_STATUS", "invalid user status")
)
