package image

import "nfxid/pkgs/errx"

const (
	CodeImageNotFound            = "IMAGE_NOT_FOUND"
	CodeFilenameRequired         = "FILENAME_REQUIRED"
	CodeOriginalFilenameRequired = "ORIGINAL_FILENAME_REQUIRED"
	CodeMimeTypeRequired         = "MIME_TYPE_REQUIRED"
	CodeSizeRequired             = "SIZE_REQUIRED"
)

var (
	ErrImageNotFound            = errx.NotFound(CodeImageNotFound, "image not found")
	ErrFilenameRequired         = errx.InvalidArg(CodeFilenameRequired, "filename is required")
	ErrOriginalFilenameRequired = errx.InvalidArg(CodeOriginalFilenameRequired, "original filename is required")
	ErrMimeTypeRequired         = errx.InvalidArg(CodeMimeTypeRequired, "mime type is required")
	ErrSizeRequired             = errx.InvalidArg(CodeSizeRequired, "size is required")
)
