package login_attempts

import "nfxid/pkgs/errx"

var (
	ErrLoginAttemptNotFound = errx.NotFound("LOGIN_ATTEMPT_NOT_FOUND", "login attempt not found")
	ErrIdentifierRequired   = errx.InvalidArg("IDENTIFIER_REQUIRED", "identifier is required")
	ErrInvalidFailureCode  = errx.InvalidArg("INVALID_FAILURE_CODE", "invalid failure code")
)
