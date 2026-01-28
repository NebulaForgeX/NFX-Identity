package user_images

import "errors"

var (
	ErrUserImageNotFound = errors.New("user image not found")
	ErrUserIDRequired    = errors.New("user id is required")
	ErrImageIDRequired   = errors.New("image id is required")
)
