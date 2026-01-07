package errors

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrUserEmailRequired     = errors.New("email is required")
	ErrUserPhoneRequired     = errors.New("phone is required")
	ErrUserUsernameRequired  = errors.New("username is required")
	ErrUserPasswordRequired  = errors.New("password is required")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrPhoneAlreadyExists    = errors.New("phone already exists")
	ErrUserInactive          = errors.New("user is inactive")
	ErrUserNotVerified       = errors.New("user is not verified")
)
