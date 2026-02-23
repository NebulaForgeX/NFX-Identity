package user_preferences

import "nfxid/pkgs/errx"

var (
	ErrUserPreferenceNotFound     = errx.NotFound("USER_PREFERENCE_NOT_FOUND", "user preference not found")
	ErrUserIDRequired             = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrUserPreferenceAlreadyExists = errx.Conflict("USER_PREFERENCE_ALREADY_EXISTS", "user preference already exists")
)
