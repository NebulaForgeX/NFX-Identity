package results

import (
	"time"

	"nfxid/modules/clients/domain/client_scopes"

	"github.com/google/uuid"
)

type ClientScopeRO struct {
	ID           uuid.UUID
	AppID        uuid.UUID
	Scope        string
	GrantedBy    *uuid.UUID
	GrantedAt    time.Time
	ExpiresAt    *time.Time
	CreatedAt    time.Time
	RevokedAt    *time.Time
	RevokedBy    *uuid.UUID
	RevokeReason *string
}

// ClientScopeMapper 将 Domain ClientScope 转换为 Application ClientScopeRO
func ClientScopeMapper(cs *client_scopes.ClientScope) ClientScopeRO {
	if cs == nil {
		return ClientScopeRO{}
	}

	return ClientScopeRO{
		ID:           cs.ID(),
		AppID:        cs.AppID(),
		Scope:        cs.Scope(),
		GrantedBy:    cs.GrantedBy(),
		GrantedAt:    cs.GrantedAt(),
		ExpiresAt:    cs.ExpiresAt(),
		CreatedAt:    cs.CreatedAt(),
		RevokedAt:    cs.RevokedAt(),
		RevokedBy:    cs.RevokedBy(),
		RevokeReason: cs.RevokeReason(),
	}
}
