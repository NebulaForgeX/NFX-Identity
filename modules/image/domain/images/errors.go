package images

import "errors"

var (
	ErrImageNotFound      = errors.New("image not found")
	ErrFilenameRequired   = errors.New("filename is required")
	ErrOriginalFilenameRequired = errors.New("original filename is required")
	ErrMimeTypeRequired   = errors.New("mime type is required")
	ErrStoragePathRequired = errors.New("storage path is required")
	ErrSizeRequired       = errors.New("size is required")
)
