package user_badges

import "errors"

var (
	ErrUserBadgeNotFound     = errors.New("user badge not found")
	ErrUserIDRequired        = errors.New("user id is required")
	ErrBadgeIDRequired       = errors.New("badge id is required")
	ErrUserBadgeAlreadyExists = errors.New("user badge already exists")
)
