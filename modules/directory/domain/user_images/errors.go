package user_images

import "nfxid/pkgs/errx"

var (
	ErrUserImageNotFound = errx.NotFound("USER_IMAGE_NOT_FOUND", "user image not found")
	ErrUserIDRequired    = errx.InvalidArg("USER_ID_REQUIRED", "user id is required")
	ErrImageIDRequired   = errx.InvalidArg("IMAGE_ID_REQUIRED", "image id is required")
)
