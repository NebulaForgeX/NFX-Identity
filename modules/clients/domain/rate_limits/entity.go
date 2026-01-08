package rate_limits

import (
	"time"

	"github.com/google/uuid"
)

type RateLimitType string

const (
	RateLimitTypeRequestsPerSecond RateLimitType = "requests_per_second"
	RateLimitTypeRequestsPerMinute RateLimitType = "requests_per_minute"
	RateLimitTypeRequestsPerHour   RateLimitType = "requests_per_hour"
	RateLimitTypeRequestsPerDay    RateLimitType = "requests_per_day"
)

type RateLimit struct {
	state RateLimitState
}

type RateLimitState struct {
	ID            uuid.UUID
	AppID         uuid.UUID
	LimitType     RateLimitType
	LimitValue    int
	WindowSeconds int
	Description   *string
	Status        string
	CreatedAt     time.Time
	CreatedBy     *uuid.UUID
	UpdatedAt     time.Time
	UpdatedBy     *uuid.UUID
}

func (rl *RateLimit) ID() uuid.UUID            { return rl.state.ID }
func (rl *RateLimit) AppID() uuid.UUID         { return rl.state.AppID }
func (rl *RateLimit) LimitType() RateLimitType { return rl.state.LimitType }
func (rl *RateLimit) LimitValue() int          { return rl.state.LimitValue }
func (rl *RateLimit) WindowSeconds() int       { return rl.state.WindowSeconds }
func (rl *RateLimit) Description() *string     { return rl.state.Description }
func (rl *RateLimit) Status() string           { return rl.state.Status }
func (rl *RateLimit) CreatedAt() time.Time     { return rl.state.CreatedAt }
func (rl *RateLimit) CreatedBy() *uuid.UUID    { return rl.state.CreatedBy }
func (rl *RateLimit) UpdatedAt() time.Time     { return rl.state.UpdatedAt }
func (rl *RateLimit) UpdatedBy() *uuid.UUID    { return rl.state.UpdatedBy }
