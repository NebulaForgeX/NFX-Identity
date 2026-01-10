package respdto

import (
	"time"

	clientCredentialAppResult "nfxid/modules/clients/application/client_credentials/results"

	"github.com/google/uuid"
)

type ClientCredentialDTO struct {
	ID           uuid.UUID  `json:"id"`
	AppID        uuid.UUID  `json:"app_id"`
	ClientID     string     `json:"client_id"`
	SecretHash   string     `json:"secret_hash"`
	HashAlg      string     `json:"hash_alg"`
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	RotatedAt    *time.Time `json:"rotated_at,omitempty"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
	LastUsedAt   *time.Time `json:"last_used_at,omitempty"`
	CreatedBy    *uuid.UUID `json:"created_by,omitempty"`
	RevokedAt    *time.Time `json:"revoked_at,omitempty"`
	RevokedBy    *uuid.UUID `json:"revoked_by,omitempty"`
	RevokeReason *string    `json:"revoke_reason,omitempty"`
}

// ClientCredentialROToDTO converts application ClientCredentialRO to response DTO
func ClientCredentialROToDTO(v *clientCredentialAppResult.ClientCredentialRO) *ClientCredentialDTO {
	if v == nil {
		return nil
	}

	return &ClientCredentialDTO{
		ID:           v.ID,
		AppID:        v.AppID,
		ClientID:     v.ClientID,
		SecretHash:   v.SecretHash,
		HashAlg:      v.HashAlg,
		Status:       string(v.Status),
		CreatedAt:    v.CreatedAt,
		RotatedAt:    v.RotatedAt,
		ExpiresAt:    v.ExpiresAt,
		LastUsedAt:   v.LastUsedAt,
		CreatedBy:    v.CreatedBy,
		RevokedAt:    v.RevokedAt,
		RevokedBy:    v.RevokedBy,
		RevokeReason: v.RevokeReason,
	}
}
