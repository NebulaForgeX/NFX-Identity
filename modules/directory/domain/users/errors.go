package users

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUsernameRequired  = errors.New("username is required")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrInvalidUserStatus = errors.New("invalid user status")
)
