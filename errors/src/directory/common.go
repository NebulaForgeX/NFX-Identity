package directory

import "nfxid/pkgs/errx"

// Shared codes and errors used by multiple directory domains.
const (
	CodeUserIDRequired  = "USER_ID_REQUIRED"
	CodeImageIDRequired = "IMAGE_ID_REQUIRED"
)

var (
	ErrUserIDRequired  = errx.InvalidArg(CodeUserIDRequired, "user id is required")
	ErrImageIDRequired = errx.InvalidArg(CodeImageIDRequired, "image id is required")
)
