package directory

import "nfxid/pkgs/errx"

const (
	CodeUserBadgeNotFound      = "USER_BADGE_NOT_FOUND"
	CodeBadgeIDRequired        = "BADGE_ID_REQUIRED"
	CodeUserBadgeAlreadyExists = "USER_BADGE_ALREADY_EXISTS"
)

var (
	ErrUserBadgeNotFound      = errx.NotFound(CodeUserBadgeNotFound, "user badge not found")
	ErrBadgeIDRequired        = errx.InvalidArg(CodeBadgeIDRequired, "badge id is required")
	ErrUserBadgeAlreadyExists = errx.Conflict(CodeUserBadgeAlreadyExists, "user badge already exists")
)
