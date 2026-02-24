package directory

import "nfxid/pkgs/errx"

const (
	CodeUserProfileNotFound      = "USER_PROFILE_NOT_FOUND"
	CodeUserProfileAlreadyExists = "USER_PROFILE_ALREADY_EXISTS"
)

var (
	ErrUserProfileNotFound      = errx.NotFound(CodeUserProfileNotFound, "user profile not found")
	ErrUserProfileAlreadyExists = errx.Conflict(CodeUserProfileAlreadyExists, "user profile already exists")
)
