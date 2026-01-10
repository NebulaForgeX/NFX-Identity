package respdto

import (
	"time"

	rateLimitAppResult "nfxid/modules/clients/application/rate_limits/results"

	"github.com/google/uuid"
)

type RateLimitDTO struct {
	ID            uuid.UUID  `json:"id"`
	AppID         uuid.UUID  `json:"app_id"`
	LimitType     string      `json:"limit_type"`
	LimitValue    int         `json:"limit_value"`
	WindowSeconds int         `json:"window_seconds"`
	Description   *string     `json:"description,omitempty"`
	Status        string      `json:"status"`
	CreatedAt     time.Time   `json:"created_at"`
	CreatedBy     *uuid.UUID  `json:"created_by,omitempty"`
	UpdatedAt     time.Time   `json:"updated_at"`
	UpdatedBy     *uuid.UUID  `json:"updated_by,omitempty"`
}

// RateLimitROToDTO converts application RateLimitRO to response DTO
func RateLimitROToDTO(v *rateLimitAppResult.RateLimitRO) *RateLimitDTO {
	if v == nil {
		return nil
	}

	return &RateLimitDTO{
		ID:            v.ID,
		AppID:         v.AppID,
		LimitType:     string(v.LimitType),
		LimitValue:    v.LimitValue,
		WindowSeconds: v.WindowSeconds,
		Description:   v.Description,
		Status:        v.Status,
		CreatedAt:     v.CreatedAt,
		CreatedBy:     v.CreatedBy,
		UpdatedAt:     v.UpdatedAt,
		UpdatedBy:     v.UpdatedBy,
	}
}
