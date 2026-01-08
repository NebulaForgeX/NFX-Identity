package user_emails

import "errors"

var (
	ErrUserEmailNotFound      = errors.New("user email not found")
	ErrUserIDRequired         = errors.New("user id is required")
	ErrEmailRequired          = errors.New("email is required")
	ErrEmailAlreadyExists     = errors.New("email already exists")
	ErrInvalidEmail           = errors.New("invalid email format")
)
