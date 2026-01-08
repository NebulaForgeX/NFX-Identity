package rate_limits

import "github.com/google/uuid"

func (rl *RateLimit) Validate() error {
	if rl.AppID() == uuid.Nil {
		return ErrAppIDRequired
	}
	if rl.LimitType() == "" {
		return ErrLimitTypeRequired
	}
	validTypes := map[RateLimitType]struct{}{
		RateLimitTypeRequestsPerSecond: {},
		RateLimitTypeRequestsPerMinute: {},
		RateLimitTypeRequestsPerHour:   {},
		RateLimitTypeRequestsPerDay:    {},
	}
	if _, ok := validTypes[rl.LimitType()]; !ok {
		return ErrInvalidRateLimitType
	}
	if rl.LimitValue() <= 0 {
		return ErrLimitValueRequired
	}
	if rl.WindowSeconds() <= 0 {
		return ErrWindowSecondsRequired
	}
	return nil
}
