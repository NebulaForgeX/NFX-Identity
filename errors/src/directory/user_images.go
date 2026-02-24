package directory

import "nfxid/pkgs/errx"

const (
	CodeUserImageNotFound = "USER_IMAGE_NOT_FOUND"
)

var (
	ErrUserImageNotFound = errx.NotFound(CodeUserImageNotFound, "user image not found")
)
