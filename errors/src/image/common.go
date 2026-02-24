package image

import "nfxid/pkgs/errx"

// Shared codes and errors used by multiple image domains.
const (
	CodeStoragePathRequired = "STORAGE_PATH_REQUIRED"
	CodeImageIDRequired     = "IMAGE_ID_REQUIRED"
)

var (
	ErrStoragePathRequired = errx.InvalidArg(CodeStoragePathRequired, "storage path is required")
	ErrImageIDRequired     = errx.InvalidArg(CodeImageIDRequired, "image id is required")
)
