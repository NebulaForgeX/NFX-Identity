package login_attempts

import "errors"

var (
	ErrLoginAttemptNotFound = errors.New("login attempt not found")
	ErrTenantIDRequired     = errors.New("tenant id is required")
	ErrIdentifierRequired   = errors.New("identifier is required")
	ErrInvalidFailureCode   = errors.New("invalid failure code")
)
