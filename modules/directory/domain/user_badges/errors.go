package user_badges

import "nfxid/pkgs/errx"

var (
	ErrUserBadgeNotFound     = errx.NotFound("USER_BADGE_NOT_FOUND", "user badge not found")
	ErrUserIDRequired        = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrBadgeIDRequired       = errx.InvalidArg("BADGE_ID_REQUIRED", "badge id is required")
	ErrUserBadgeAlreadyExists = errx.Conflict("USER_BADGE_ALREADY_EXISTS", "user badge already exists")
)
