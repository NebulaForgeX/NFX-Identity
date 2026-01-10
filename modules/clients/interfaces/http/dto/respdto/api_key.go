package respdto

import (
	"time"

	apiKeyAppResult "nfxid/modules/clients/application/api_keys/results"

	"github.com/google/uuid"
)

type APIKeyDTO struct {
	ID           uuid.UUID              `json:"id"`
	KeyID        string                 `json:"key_id"`
	AppID        uuid.UUID              `json:"app_id"`
	KeyHash      string                 `json:"key_hash"`
	HashAlg      string                 `json:"hash_alg"`
	Name         string                 `json:"name"`
	Status       string                 `json:"status"`
	ExpiresAt    *time.Time             `json:"expires_at,omitempty"`
	CreatedAt    time.Time              `json:"created_at"`
	RevokedAt    *time.Time             `json:"revoked_at,omitempty"`
	RevokedBy    *uuid.UUID             `json:"revoked_by,omitempty"`
	RevokeReason *string                `json:"revoke_reason,omitempty"`
	LastUsedAt   *time.Time             `json:"last_used_at,omitempty"`
	CreatedBy    *uuid.UUID             `json:"created_by,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

// APIKeyROToDTO converts application APIKeyRO to response DTO
func APIKeyROToDTO(v *apiKeyAppResult.APIKeyRO) *APIKeyDTO {
	if v == nil {
		return nil
	}

	return &APIKeyDTO{
		ID:           v.ID,
		KeyID:        v.KeyID,
		AppID:        v.AppID,
		KeyHash:      v.KeyHash,
		HashAlg:      v.HashAlg,
		Name:         v.Name,
		Status:       string(v.Status),
		ExpiresAt:    v.ExpiresAt,
		CreatedAt:    v.CreatedAt,
		RevokedAt:    v.RevokedAt,
		RevokedBy:    v.RevokedBy,
		RevokeReason: v.RevokeReason,
		LastUsedAt:   v.LastUsedAt,
		CreatedBy:    v.CreatedBy,
		Metadata:     v.Metadata,
	}
}
