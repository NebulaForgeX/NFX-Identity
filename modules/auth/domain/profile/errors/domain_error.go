package errors

import "errors"

var (
	ErrProfileNotFound      = errors.New("profile not found")
	ErrProfileUserIDRequired = errors.New("user_id is required")
	ErrProfileAlreadyExists  = errors.New("profile already exists for this user")
)

