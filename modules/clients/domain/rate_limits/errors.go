package rate_limits

import "errors"

var (
	ErrRateLimitNotFound      = errors.New("rate limit not found")
	ErrAppIDRequired          = errors.New("app id is required")
	ErrLimitTypeRequired      = errors.New("limit type is required")
	ErrLimitValueRequired     = errors.New("limit value is required")
	ErrWindowSecondsRequired  = errors.New("window seconds is required")
	ErrRateLimitAlreadyExists = errors.New("rate limit already exists")
	ErrInvalidRateLimitType   = errors.New("invalid rate limit type")
	ErrInvalidStatus          = errors.New("invalid status")
)
