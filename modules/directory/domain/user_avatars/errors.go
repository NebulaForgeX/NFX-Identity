package user_avatars

import "nfxid/pkgs/errx"

var (
	ErrUserAvatarNotFound = errx.NotFound("USER_AVATAR_NOT_FOUND", "user avatar not found")
	ErrUserIDRequired     = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrImageIDRequired    = errx.InvalidArg("IMAGE_ID_REQUIRED", "image id is required")
)
