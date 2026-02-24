package image

import "nfxid/pkgs/errx"

const (
	CodeImageVariantNotFound      = "IMAGE_VARIANT_NOT_FOUND"
	CodeVariantKeyRequired        = "VARIANT_KEY_REQUIRED"
	CodeImageVariantAlreadyExists = "IMAGE_VARIANT_ALREADY_EXISTS"
)

var (
	ErrImageVariantNotFound      = errx.NotFound(CodeImageVariantNotFound, "image variant not found")
	ErrVariantKeyRequired        = errx.InvalidArg(CodeVariantKeyRequired, "variant key is required")
	ErrImageVariantAlreadyExists = errx.Conflict(CodeImageVariantAlreadyExists, "image variant already exists")
)
