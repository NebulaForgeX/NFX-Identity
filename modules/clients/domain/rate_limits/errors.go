package rate_limits

import "nfxid/pkgs/errx"

var (
	ErrRateLimitNotFound      = errx.NotFound("RATE_LIMIT_NOT_FOUND", "rate limit not found")
	ErrAppIDRequired          = errx.InvalidArg("APP_ID_REQUIRED", "app id is required")
	ErrLimitTypeRequired      = errx.InvalidArg("LIMIT_TYPE_REQUIRED", "limit type is required")
	ErrLimitValueRequired     = errx.InvalidArg("LIMIT_VALUE_REQUIRED", "limit value is required")
	ErrWindowSecondsRequired  = errx.InvalidArg("WINDOW_SECONDS_REQUIRED", "window seconds is required")
	ErrRateLimitAlreadyExists  = errx.Conflict("RATE_LIMIT_ALREADY_EXISTS", "rate limit already exists")
	ErrInvalidRateLimitType   = errx.InvalidArg("INVALID_RATE_LIMIT_TYPE", "invalid rate limit type")
	ErrInvalidStatus          = errx.InvalidArg("INVALID_STATUS", "invalid status")
)
