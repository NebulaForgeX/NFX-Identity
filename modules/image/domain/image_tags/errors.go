package image_tags

import "errors"

var (
	ErrImageTagNotFound      = errors.New("image tag not found")
	ErrImageIDRequired       = errors.New("image id is required")
	ErrTagRequired           = errors.New("tag is required")
	ErrImageTagAlreadyExists = errors.New("image tag already exists")
)
