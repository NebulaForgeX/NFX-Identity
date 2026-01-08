package rate_limits

import (
	"time"

	"github.com/google/uuid"
)

func (rl *RateLimit) UpdateLimit(limitValue, windowSeconds int, description *string, updatedBy *uuid.UUID) error {
	if limitValue <= 0 {
		return ErrLimitValueRequired
	}
	if windowSeconds <= 0 {
		return ErrWindowSecondsRequired
	}

	rl.state.LimitValue = limitValue
	rl.state.WindowSeconds = windowSeconds
	if description != nil {
		rl.state.Description = description
	}
	rl.state.UpdatedBy = updatedBy
	rl.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (rl *RateLimit) UpdateStatus(status string, updatedBy *uuid.UUID) error {
	validStatuses := map[string]struct{}{
		"active":   {},
		"disabled": {},
	}
	if _, ok := validStatuses[status]; !ok {
		return ErrInvalidStatus
	}

	rl.state.Status = status
	rl.state.UpdatedBy = updatedBy
	rl.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (rl *RateLimit) IsActive() bool {
	return rl.Status() == "active"
}
