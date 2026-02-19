package reqdto

import (
	apiKeyAppCommands "nfxid/modules/clients/application/api_keys/commands"

	"github.com/google/uuid"
)

type APIKeyCreateRequestDTO struct {
	KeyID     string                 `json:"key_id" validate:"required"`
	AppID     uuid.UUID              `json:"app_id" validate:"required"`
	KeyHash   string                 `json:"key_hash" validate:"required"`
	HashAlg   string                 `json:"hash_alg" validate:"required"`
	Name      string                 `json:"name" validate:"required"`
	ExpiresAt *string                `json:"expires_at,omitempty"`
	CreatedBy *uuid.UUID             `json:"created_by,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type APIKeyByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type APIKeyDeleteRequestDTO struct {
	KeyID string `uri:"key_id" validate:"required"`
}

func (r *APIKeyCreateRequestDTO) ToCreateCmd() apiKeyAppCommands.CreateAPIKeyCmd {
	return apiKeyAppCommands.CreateAPIKeyCmd{
		KeyID:     r.KeyID,
		AppID:     r.AppID,
		KeyHash:   r.KeyHash,
		HashAlg:   r.HashAlg,
		Name:      r.Name,
		ExpiresAt: r.ExpiresAt,
		CreatedBy: r.CreatedBy,
		Metadata:  r.Metadata,
	}
}
