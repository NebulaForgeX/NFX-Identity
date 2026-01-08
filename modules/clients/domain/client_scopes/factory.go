package client_scopes

import (
	"time"

	"github.com/google/uuid"
)

type NewClientScopeParams struct {
	AppID     uuid.UUID
	Scope     string
	GrantedBy *uuid.UUID
	ExpiresAt *time.Time
}

func NewClientScope(p NewClientScopeParams) (*ClientScope, error) {
	if err := validateClientScopeParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewClientScopeFromState(ClientScopeState{
		ID:        id,
		AppID:     p.AppID,
		Scope:     p.Scope,
		GrantedBy: p.GrantedBy,
		GrantedAt: now,
		ExpiresAt: p.ExpiresAt,
		CreatedAt: now,
	}), nil
}

func NewClientScopeFromState(st ClientScopeState) *ClientScope {
	return &ClientScope{state: st}
}

func validateClientScopeParams(p NewClientScopeParams) error {
	if p.AppID == uuid.Nil {
		return ErrAppIDRequired
	}
	if p.Scope == "" {
		return ErrScopeRequired
	}
	return nil
}
