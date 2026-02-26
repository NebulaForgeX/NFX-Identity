package rate_limits

import (
	clientsErr "nfxidentity/errors/src/clients"

	"github.com/google/uuid"
)

func (rl *RateLimit) Validate() error {
	if rl.AppID() == uuid.Nil {
		return clientsErr.ErrAppIDRequired
	}
	if rl.LimitType() == "" {
		return clientsErr.ErrLimitTypeRequired
	}
	validTypes := map[RateLimitType]struct{}{
		RateLimitTypeRequestsPerSecond: {},
		RateLimitTypeRequestsPerMinute: {},
		RateLimitTypeRequestsPerHour:   {},
		RateLimitTypeRequestsPerDay:    {},
	}
	if _, ok := validTypes[rl.LimitType()]; !ok {
		return clientsErr.ErrInvalidRateLimitType
	}
	if rl.LimitValue() <= 0 {
		return clientsErr.ErrLimitValueRequired
	}
	if rl.WindowSeconds() <= 0 {
		return clientsErr.ErrWindowSecondsRequired
	}
	return nil
}
