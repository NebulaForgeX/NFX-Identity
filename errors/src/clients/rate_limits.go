package clients

import "nfxid/pkgs/errx"

const (
	CodeRateLimitNotFound      = "RATE_LIMIT_NOT_FOUND"
	CodeLimitTypeRequired      = "LIMIT_TYPE_REQUIRED"
	CodeLimitValueRequired     = "LIMIT_VALUE_REQUIRED"
	CodeWindowSecondsRequired  = "WINDOW_SECONDS_REQUIRED"
	CodeRateLimitAlreadyExists = "RATE_LIMIT_ALREADY_EXISTS"
	CodeInvalidRateLimitType   = "INVALID_RATE_LIMIT_TYPE"
	CodeInvalidStatus          = "INVALID_STATUS"
)

var (
	ErrRateLimitNotFound      = errx.NotFound(CodeRateLimitNotFound, "rate limit not found")
	ErrLimitTypeRequired      = errx.InvalidArg(CodeLimitTypeRequired, "limit type is required")
	ErrLimitValueRequired     = errx.InvalidArg(CodeLimitValueRequired, "limit value is required")
	ErrWindowSecondsRequired  = errx.InvalidArg(CodeWindowSecondsRequired, "window seconds is required")
	ErrRateLimitAlreadyExists = errx.Conflict(CodeRateLimitAlreadyExists, "rate limit already exists")
	ErrInvalidRateLimitType   = errx.InvalidArg(CodeInvalidRateLimitType, "invalid rate limit type")
	ErrInvalidStatus          = errx.InvalidArg(CodeInvalidStatus, "invalid status")
)
