package image_types

import "errors"

var (
	ErrImageTypeNotFound    = errors.New("image type not found")
	ErrKeyRequired          = errors.New("key is required")
	ErrKeyAlreadyExists     = errors.New("key already exists")
	ErrCannotDeleteSystemType = errors.New("cannot delete system type")
)
