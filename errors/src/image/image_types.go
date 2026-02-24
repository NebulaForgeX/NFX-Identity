package image

import "nfxid/pkgs/errx"

const (
	CodeImageTypeNotFound      = "IMAGE_TYPE_NOT_FOUND"
	CodeKeyRequired            = "KEY_REQUIRED"
	CodeKeyAlreadyExists       = "KEY_ALREADY_EXISTS"
	CodeCannotDeleteSystemType = "CANNOT_DELETE_SYSTEM_TYPE"
)

var (
	ErrImageTypeNotFound      = errx.NotFound(CodeImageTypeNotFound, "image type not found")
	ErrKeyRequired            = errx.InvalidArg(CodeKeyRequired, "key is required")
	ErrKeyAlreadyExists       = errx.Conflict(CodeKeyAlreadyExists, "key already exists")
	ErrCannotDeleteSystemType = errx.FailedPrecond(CodeCannotDeleteSystemType, "cannot delete system type")
)
