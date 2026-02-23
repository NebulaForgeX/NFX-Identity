package image_tags

import "nfxid/pkgs/errx"

var (
	ErrImageTagNotFound      = errx.NotFound("IMAGE_TAG_NOT_FOUND", "image tag not found")
	ErrImageIDRequired       = errx.InvalidArg("IMAGE_ID_REQUIRED", "image id is required")
	ErrTagRequired           = errx.InvalidArg("TAG_REQUIRED", "tag is required")
	ErrImageTagAlreadyExists = errx.Conflict("IMAGE_TAG_ALREADY_EXISTS", "image tag already exists")
)
