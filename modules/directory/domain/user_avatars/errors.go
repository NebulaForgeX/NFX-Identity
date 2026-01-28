package user_avatars

import "errors"

var (
	ErrUserAvatarNotFound = errors.New("user avatar not found")
	ErrUserIDRequired     = errors.New("user id is required")
	ErrImageIDRequired    = errors.New("image id is required")
)
