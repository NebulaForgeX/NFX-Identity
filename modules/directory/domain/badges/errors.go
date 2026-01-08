package badges

import "errors"

var (
	ErrBadgeNotFound      = errors.New("badge not found")
	ErrNameRequired       = errors.New("name is required")
	ErrNameAlreadyExists  = errors.New("name already exists")
)
