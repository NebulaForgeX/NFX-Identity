package image

import "nfxid/pkgs/errx"

const (
	CodeImageTagNotFound      = "IMAGE_TAG_NOT_FOUND"
	CodeTagRequired           = "TAG_REQUIRED"
	CodeImageTagAlreadyExists = "IMAGE_TAG_ALREADY_EXISTS"
)

var (
	ErrImageTagNotFound      = errx.NotFound(CodeImageTagNotFound, "image tag not found")
	ErrTagRequired           = errx.InvalidArg(CodeTagRequired, "tag is required")
	ErrImageTagAlreadyExists = errx.Conflict(CodeImageTagAlreadyExists, "image tag already exists")
)
