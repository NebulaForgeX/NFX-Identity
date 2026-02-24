package directory

import "nfxid/pkgs/errx"

const (
	CodeUserAvatarNotFound = "USER_AVATAR_NOT_FOUND"
)

var (
	ErrUserAvatarNotFound = errx.NotFound(CodeUserAvatarNotFound, "user avatar not found")
)
