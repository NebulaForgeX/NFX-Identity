package errors

import "errors"

var (
	ErrProfileBadgeNotFound     = errors.New("profile badge not found")
	ErrProfileBadgeAlreadyExists = errors.New("profile badge already exists")
	ErrProfileIDRequired       = errors.New("profile ID is required for profile badge")
	ErrBadgeIDRequired          = errors.New("badge ID is required for profile badge")
)

