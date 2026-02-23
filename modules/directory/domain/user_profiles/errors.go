package user_profiles

import "nfxid/pkgs/errx"

var (
	ErrUserProfileNotFound     = errx.NotFound("USER_PROFILE_NOT_FOUND", "user profile not found")
	ErrUserIDRequired          = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrUserProfileAlreadyExists = errx.Conflict("USER_PROFILE_ALREADY_EXISTS", "user profile already exists")
)
