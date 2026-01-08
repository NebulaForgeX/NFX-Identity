package image_variants

import "errors"

var (
	ErrImageVariantNotFound      = errors.New("image variant not found")
	ErrImageIDRequired           = errors.New("image id is required")
	ErrVariantKeyRequired        = errors.New("variant key is required")
	ErrStoragePathRequired       = errors.New("storage path is required")
	ErrImageVariantAlreadyExists = errors.New("image variant already exists")
)
