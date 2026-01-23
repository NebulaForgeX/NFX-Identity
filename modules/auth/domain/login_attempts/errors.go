package login_attempts

import "errors"

var (
	ErrLoginAttemptNotFound = errors.New("login attempt not found")
	ErrIdentifierRequired   = errors.New("identifier is required")
	ErrInvalidFailureCode   = errors.New("invalid failure code")
)
