package respdto

import (
	"time"

	ipAllowlistAppResult "nfxid/modules/clients/application/ip_allowlist/results"

	"github.com/google/uuid"
)

type IPAllowlistDTO struct {
	ID          uuid.UUID  `json:"id"`
	RuleID      string     `json:"rule_id"`
	AppID       uuid.UUID  `json:"app_id"`
	CIDR        string     `json:"cidr"`
	Description *string    `json:"description,omitempty"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   *uuid.UUID `json:"created_by,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UpdatedBy   *uuid.UUID `json:"updated_by,omitempty"`
	RevokedAt   *time.Time `json:"revoked_at,omitempty"`
	RevokedBy   *uuid.UUID `json:"revoked_by,omitempty"`
	RevokeReason *string   `json:"revoke_reason,omitempty"`
}

// IPAllowlistROToDTO converts application IPAllowlistRO to response DTO
func IPAllowlistROToDTO(v *ipAllowlistAppResult.IPAllowlistRO) *IPAllowlistDTO {
	if v == nil {
		return nil
	}

	return &IPAllowlistDTO{
		ID:          v.ID,
		RuleID:      v.RuleID,
		AppID:       v.AppID,
		CIDR:        v.CIDR,
		Description: v.Description,
		Status:      string(v.Status),
		CreatedAt:   v.CreatedAt,
		CreatedBy:   v.CreatedBy,
		UpdatedAt:   v.UpdatedAt,
		UpdatedBy:   v.UpdatedBy,
		RevokedAt:   v.RevokedAt,
		RevokedBy:   v.RevokedBy,
		RevokeReason: v.RevokeReason,
	}
}
