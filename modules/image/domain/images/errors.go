package images

import "nfxid/pkgs/errx"

var (
	ErrImageNotFound           = errx.NotFound("IMAGE_NOT_FOUND", "image not found")
	ErrFilenameRequired        = errx.InvalidArg("FILENAME_REQUIRED", "filename is required")
	ErrOriginalFilenameRequired = errx.InvalidArg("ORIGINAL_FILENAME_REQUIRED", "original filename is required")
	ErrMimeTypeRequired         = errx.InvalidArg("MIME_TYPE_REQUIRED", "mime type is required")
	ErrStoragePathRequired     = errx.InvalidArg("STORAGE_PATH_REQUIRED", "storage path is required")
	ErrSizeRequired            = errx.InvalidArg("SIZE_REQUIRED", "size is required")
)
