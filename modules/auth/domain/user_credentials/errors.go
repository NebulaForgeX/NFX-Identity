package user_credentials

import "errors"

var (
	ErrUserCredentialNotFound    = errors.New("user credential not found")
	ErrUserIDRequired            = errors.New("user id is required")
	ErrCredentialTypeRequired    = errors.New("credential type is required")
	ErrInvalidCredentialType     = errors.New("invalid credential type")
	ErrInvalidCredentialStatus   = errors.New("invalid credential status")
	ErrPasswordHashRequired      = errors.New("password hash is required")
	ErrUserCredentialAlreadyExists = errors.New("user credential already exists")
)
