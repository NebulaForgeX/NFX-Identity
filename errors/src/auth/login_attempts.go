package auth

import "nfxid/pkgs/errx"

const (
	CodeLoginAttemptNotFound = "LOGIN_ATTEMPT_NOT_FOUND"
	CodeIdentifierRequired   = "IDENTIFIER_REQUIRED"
	CodeInvalidFailureCode   = "INVALID_FAILURE_CODE"
)

var (
	ErrLoginAttemptNotFound = errx.NotFound(CodeLoginAttemptNotFound, "login attempt not found")
	ErrIdentifierRequired   = errx.InvalidArg(CodeIdentifierRequired, "identifier is required")
	ErrInvalidFailureCode   = errx.InvalidArg(CodeInvalidFailureCode, "invalid failure code")
)
