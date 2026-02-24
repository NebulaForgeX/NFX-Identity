package directory

import "nfxid/pkgs/errx"

const (
	CodeUserPreferenceNotFound      = "USER_PREFERENCE_NOT_FOUND"
	CodeUserPreferenceAlreadyExists = "USER_PREFERENCE_ALREADY_EXISTS"
)

var (
	ErrUserPreferenceNotFound      = errx.NotFound(CodeUserPreferenceNotFound, "user preference not found")
	ErrUserPreferenceAlreadyExists = errx.Conflict(CodeUserPreferenceAlreadyExists, "user preference already exists")
)
