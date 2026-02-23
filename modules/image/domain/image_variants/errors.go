package image_variants

import "nfxid/pkgs/errx"

var (
	ErrImageVariantNotFound      = errx.NotFound("IMAGE_VARIANT_NOT_FOUND", "image variant not found")
	ErrImageIDRequired           = errx.InvalidArg("IMAGE_ID_REQUIRED", "image id is required")
	ErrVariantKeyRequired        = errx.InvalidArg("VARIANT_KEY_REQUIRED", "variant key is required")
	ErrStoragePathRequired       = errx.InvalidArg("STORAGE_PATH_REQUIRED", "storage path is required")
	ErrImageVariantAlreadyExists = errx.Conflict("IMAGE_VARIANT_ALREADY_EXISTS", "image variant already exists")
)
