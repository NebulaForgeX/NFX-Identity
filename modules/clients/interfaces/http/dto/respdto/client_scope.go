package respdto

import (
	"time"

	clientScopeAppResult "nfxid/modules/clients/application/client_scopes/results"

	"github.com/google/uuid"
)

type ClientScopeDTO struct {
	ID          uuid.UUID  `json:"id"`
	AppID       uuid.UUID  `json:"app_id"`
	Scope       string     `json:"scope"`
	GrantedBy   *uuid.UUID `json:"granted_by,omitempty"`
	GrantedAt   time.Time  `json:"granted_at"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	RevokedAt   *time.Time `json:"revoked_at,omitempty"`
	RevokedBy   *uuid.UUID `json:"revoked_by,omitempty"`
	RevokeReason *string   `json:"revoke_reason,omitempty"`
}

// ClientScopeROToDTO converts application ClientScopeRO to response DTO
func ClientScopeROToDTO(v *clientScopeAppResult.ClientScopeRO) *ClientScopeDTO {
	if v == nil {
		return nil
	}

	return &ClientScopeDTO{
		ID:          v.ID,
		AppID:       v.AppID,
		Scope:       v.Scope,
		GrantedBy:   v.GrantedBy,
		GrantedAt:   v.GrantedAt,
		ExpiresAt:   v.ExpiresAt,
		CreatedAt:   v.CreatedAt,
		RevokedAt:   v.RevokedAt,
		RevokedBy:   v.RevokedBy,
		RevokeReason: v.RevokeReason,
	}
}
