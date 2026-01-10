package results

import (
	"time"

	"nfxid/modules/clients/domain/rate_limits"

	"github.com/google/uuid"
)

type RateLimitRO struct {
	ID            uuid.UUID
	AppID         uuid.UUID
	LimitType     rate_limits.RateLimitType
	LimitValue    int
	WindowSeconds int
	Description   *string
	Status        string
	CreatedAt     time.Time
	CreatedBy     *uuid.UUID
	UpdatedAt     time.Time
	UpdatedBy     *uuid.UUID
}

// RateLimitMapper 将 Domain RateLimit 转换为 Application RateLimitRO
func RateLimitMapper(rl *rate_limits.RateLimit) RateLimitRO {
	if rl == nil {
		return RateLimitRO{}
	}

	return RateLimitRO{
		ID:            rl.ID(),
		AppID:         rl.AppID(),
		LimitType:     rl.LimitType(),
		LimitValue:    rl.LimitValue(),
		WindowSeconds: rl.WindowSeconds(),
		Description:   rl.Description(),
		Status:        rl.Status(),
		CreatedAt:     rl.CreatedAt(),
		CreatedBy:     rl.CreatedBy(),
		UpdatedAt:     rl.UpdatedAt(),
		UpdatedBy:     rl.UpdatedBy(),
	}
}
