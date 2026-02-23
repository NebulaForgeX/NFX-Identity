package image_types

import "nfxid/pkgs/errx"

var (
	ErrImageTypeNotFound     = errx.NotFound("IMAGE_TYPE_NOT_FOUND", "image type not found")
	ErrKeyRequired           = errx.InvalidArg("KEY_REQUIRED", "key is required")
	ErrKeyAlreadyExists      = errx.Conflict("KEY_ALREADY_EXISTS", "key already exists")
	ErrCannotDeleteSystemType = errx.FailedPrecond("CANNOT_DELETE_SYSTEM_TYPE", "cannot delete system type")
)
