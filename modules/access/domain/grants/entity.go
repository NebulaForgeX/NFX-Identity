package grants

import (
	"time"

	"github.com/google/uuid"
)

type SubjectType string

const (
	SubjectTypeUser   SubjectType = "USER"
	SubjectTypeClient SubjectType = "CLIENT"
)

type GrantType string

const (
	GrantTypeRole       GrantType = "ROLE"
	GrantTypePermission GrantType = "PERMISSION"
)

type GrantEffect string

const (
	GrantEffectAllow GrantEffect = "ALLOW"
	GrantEffectDeny  GrantEffect = "DENY"
)

type Grant struct {
	state GrantState
}

type GrantState struct {
	ID           uuid.UUID
	SubjectType  SubjectType
	SubjectID    uuid.UUID
	GrantType    GrantType
	GrantRefID   uuid.UUID
	TenantID     *uuid.UUID
	AppID        *uuid.UUID
	ResourceType *string
	ResourceID   *uuid.UUID
	Effect       GrantEffect
	ExpiresAt    *time.Time
	CreatedAt    time.Time
	CreatedBy    *uuid.UUID
	RevokedAt    *time.Time
	RevokedBy    *uuid.UUID
	RevokeReason *string
}

func (g *Grant) ID() uuid.UUID          { return g.state.ID }
func (g *Grant) SubjectType() SubjectType { return g.state.SubjectType }
func (g *Grant) SubjectID() uuid.UUID   { return g.state.SubjectID }
func (g *Grant) GrantType() GrantType   { return g.state.GrantType }
func (g *Grant) GrantRefID() uuid.UUID  { return g.state.GrantRefID }
func (g *Grant) TenantID() *uuid.UUID   { return g.state.TenantID }
func (g *Grant) AppID() *uuid.UUID      { return g.state.AppID }
func (g *Grant) ResourceType() *string  { return g.state.ResourceType }
func (g *Grant) ResourceID() *uuid.UUID { return g.state.ResourceID }
func (g *Grant) Effect() GrantEffect    { return g.state.Effect }
func (g *Grant) ExpiresAt() *time.Time  { return g.state.ExpiresAt }
func (g *Grant) CreatedAt() time.Time   { return g.state.CreatedAt }
func (g *Grant) CreatedBy() *uuid.UUID  { return g.state.CreatedBy }
func (g *Grant) RevokedAt() *time.Time  { return g.state.RevokedAt }
func (g *Grant) RevokedBy() *uuid.UUID  { return g.state.RevokedBy }
func (g *Grant) RevokeReason() *string  { return g.state.RevokeReason }
