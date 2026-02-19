package reqdto

import (
	clientCredentialAppCommands "nfxid/modules/clients/application/client_credentials/commands"

	"github.com/google/uuid"
)

type ClientCredentialCreateRequestDTO struct {
	AppID      uuid.UUID  `json:"app_id" validate:"required"`
	ClientID   string     `json:"client_id" validate:"required"`
	SecretHash string     `json:"secret_hash" validate:"required"`
	HashAlg    string     `json:"hash_alg" validate:"required"`
	ExpiresAt  *string    `json:"expires_at,omitempty"`
	CreatedBy  *uuid.UUID `json:"created_by,omitempty"`
}

type ClientCredentialByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type ClientCredentialDeleteRequestDTO struct {
	ClientID string `uri:"client_id" validate:"required"`
}

func (r *ClientCredentialCreateRequestDTO) ToCreateCmd() clientCredentialAppCommands.CreateClientCredentialCmd {
	return clientCredentialAppCommands.CreateClientCredentialCmd{
		AppID:      r.AppID,
		ClientID:   r.ClientID,
		SecretHash: r.SecretHash,
		HashAlg:    r.HashAlg,
		ExpiresAt:  r.ExpiresAt,
		CreatedBy:  r.CreatedBy,
	}
}
