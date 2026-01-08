package client_scopes

import (
	"time"

	"github.com/google/uuid"
)

type ClientScope struct {
	state ClientScopeState
}

type ClientScopeState struct {
	ID          uuid.UUID
	AppID       uuid.UUID
	Scope       string
	GrantedBy   *uuid.UUID
	GrantedAt   time.Time
	ExpiresAt   *time.Time
	CreatedAt   time.Time
	RevokedAt   *time.Time
	RevokedBy   *uuid.UUID
	RevokeReason *string
}

func (cs *ClientScope) ID() uuid.UUID             { return cs.state.ID }
func (cs *ClientScope) AppID() uuid.UUID           { return cs.state.AppID }
func (cs *ClientScope) Scope() string              { return cs.state.Scope }
func (cs *ClientScope) GrantedBy() *uuid.UUID      { return cs.state.GrantedBy }
func (cs *ClientScope) GrantedAt() time.Time       { return cs.state.GrantedAt }
func (cs *ClientScope) ExpiresAt() *time.Time      { return cs.state.ExpiresAt }
func (cs *ClientScope) CreatedAt() time.Time       { return cs.state.CreatedAt }
func (cs *ClientScope) RevokedAt() *time.Time      { return cs.state.RevokedAt }
func (cs *ClientScope) RevokedBy() *uuid.UUID      { return cs.state.RevokedBy }
func (cs *ClientScope) RevokeReason() *string      { return cs.state.RevokeReason }
