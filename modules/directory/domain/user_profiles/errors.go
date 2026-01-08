package user_profiles

import "errors"

var (
	ErrUserProfileNotFound   = errors.New("user profile not found")
	ErrUserIDRequired        = errors.New("user id is required")
	ErrUserProfileAlreadyExists = errors.New("user profile already exists")
)
