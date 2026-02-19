package reqdto

import (
	rateLimitAppCommands "nfxid/modules/clients/application/rate_limits/commands"
	rateLimitDomain "nfxid/modules/clients/domain/rate_limits"

	"github.com/google/uuid"
)

type RateLimitCreateRequestDTO struct {
	AppID         uuid.UUID  `json:"app_id" validate:"required"`
	LimitType     string     `json:"limit_type" validate:"required"`
	LimitValue    int        `json:"limit_value" validate:"required"`
	WindowSeconds int        `json:"window_seconds" validate:"required"`
	Description   *string    `json:"description,omitempty"`
	Status        string     `json:"status,omitempty"`
	CreatedBy     *uuid.UUID `json:"created_by,omitempty"`
}

type RateLimitByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

func (r *RateLimitCreateRequestDTO) ToCreateCmd() rateLimitAppCommands.CreateRateLimitCmd {
	cmd := rateLimitAppCommands.CreateRateLimitCmd{
		AppID:         r.AppID,
		LimitValue:    r.LimitValue,
		WindowSeconds: r.WindowSeconds,
		Description:   r.Description,
		Status:        r.Status,
		CreatedBy:     r.CreatedBy,
	}

	if r.LimitType != "" {
		cmd.LimitType = rateLimitDomain.RateLimitType(r.LimitType)
	}

	return cmd
}
