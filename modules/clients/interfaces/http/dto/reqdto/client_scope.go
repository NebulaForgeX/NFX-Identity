package reqdto

import (
	clientScopeAppCommands "nfxid/modules/clients/application/client_scopes/commands"

	"github.com/google/uuid"
)

type ClientScopeCreateRequestDTO struct {
	AppID     uuid.UUID  `json:"app_id" validate:"required"`
	Scope     string     `json:"scope" validate:"required"`
	GrantedBy *uuid.UUID `json:"granted_by,omitempty"`
	ExpiresAt *string    `json:"expires_at,omitempty"`
}

type ClientScopeByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

func (r *ClientScopeCreateRequestDTO) ToCreateCmd() clientScopeAppCommands.CreateClientScopeCmd {
	return clientScopeAppCommands.CreateClientScopeCmd{
		AppID:     r.AppID,
		Scope:     r.Scope,
		GrantedBy: r.GrantedBy,
		ExpiresAt: r.ExpiresAt,
	}
}
