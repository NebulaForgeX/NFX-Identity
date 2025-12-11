package auth

import "errors"

var (
	ErrLoginIdentifierRequired = errors.New("login identifier is required")
	ErrInvalidCredentials      = errors.New("invalid credentials")
	ErrUserInactive            = errors.New("user is inactive")
	ErrEmailCodeNotImplemented = errors.New("email code login not implemented yet")
)

