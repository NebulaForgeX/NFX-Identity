package rate_limits

import (
	"time"

	"github.com/google/uuid"
)

type NewRateLimitParams struct {
	AppID         uuid.UUID
	LimitType     RateLimitType
	LimitValue    int
	WindowSeconds int
	Description   *string
	Status        string
	CreatedBy     *uuid.UUID
}

func NewRateLimit(p NewRateLimitParams) (*RateLimit, error) {
	if err := validateRateLimitParams(p); err != nil {
		return nil, err
	}

	limitType := p.LimitType
	if limitType == "" {
		limitType = RateLimitTypeRequestsPerMinute
	}

	status := p.Status
	if status == "" {
		status = "active"
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewRateLimitFromState(RateLimitState{
		ID:            id,
		AppID:         p.AppID,
		LimitType:     limitType,
		LimitValue:    p.LimitValue,
		WindowSeconds: p.WindowSeconds,
		Description:   p.Description,
		Status:        status,
		CreatedAt:     now,
		UpdatedAt:     now,
		CreatedBy:     p.CreatedBy,
	}), nil
}

func NewRateLimitFromState(st RateLimitState) *RateLimit {
	return &RateLimit{state: st}
}

func validateRateLimitParams(p NewRateLimitParams) error {
	if p.AppID == uuid.Nil {
		return ErrAppIDRequired
	}
	if p.LimitType == "" {
		return ErrLimitTypeRequired
	}
	validTypes := map[RateLimitType]struct{}{
		RateLimitTypeRequestsPerSecond: {},
		RateLimitTypeRequestsPerMinute: {},
		RateLimitTypeRequestsPerHour:   {},
		RateLimitTypeRequestsPerDay:    {},
	}
	if _, ok := validTypes[p.LimitType]; !ok {
		return ErrInvalidRateLimitType
	}
	if p.LimitValue <= 0 {
		return ErrLimitValueRequired
	}
	if p.WindowSeconds <= 0 {
		return ErrWindowSecondsRequired
	}
	return nil
}
