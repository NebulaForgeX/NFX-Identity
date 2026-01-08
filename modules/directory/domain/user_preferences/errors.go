package user_preferences

import "errors"

var (
	ErrUserPreferenceNotFound   = errors.New("user preference not found")
	ErrUserIDRequired           = errors.New("user id is required")
	ErrUserPreferenceAlreadyExists = errors.New("user preference already exists")
)
