package invitations

import (
	"time"

	"github.com/google/uuid"
)

type InvitationStatus string

const (
	InvitationStatusPending InvitationStatus = "PENDING"
	InvitationStatusAccepted InvitationStatus = "ACCEPTED"
	InvitationStatusExpired InvitationStatus = "EXPIRED"
	InvitationStatusRevoked InvitationStatus = "REVOKED"
)

type Invitation struct {
	state InvitationState
}

type InvitationState struct {
	ID              uuid.UUID
	InviteID        string
	TenantID        uuid.UUID
	Email           string
	TokenHash       string
	ExpiresAt       time.Time
	Status          InvitationStatus
	InvitedBy       uuid.UUID
	InvitedAt       time.Time
	AcceptedByUserID *uuid.UUID
	AcceptedAt      *time.Time
	RevokedBy       *uuid.UUID
	RevokedAt       *time.Time
	RevokeReason    *string
	RoleIDs         []uuid.UUID
	Metadata        map[string]interface{}
}

func (i *Invitation) ID() uuid.UUID                   { return i.state.ID }
func (i *Invitation) InviteID() string                { return i.state.InviteID }
func (i *Invitation) TenantID() uuid.UUID             { return i.state.TenantID }
func (i *Invitation) Email() string                   { return i.state.Email }
func (i *Invitation) TokenHash() string               { return i.state.TokenHash }
func (i *Invitation) ExpiresAt() time.Time            { return i.state.ExpiresAt }
func (i *Invitation) Status() InvitationStatus        { return i.state.Status }
func (i *Invitation) InvitedBy() uuid.UUID            { return i.state.InvitedBy }
func (i *Invitation) InvitedAt() time.Time            { return i.state.InvitedAt }
func (i *Invitation) AcceptedByUserID() *uuid.UUID    { return i.state.AcceptedByUserID }
func (i *Invitation) AcceptedAt() *time.Time          { return i.state.AcceptedAt }
func (i *Invitation) RevokedBy() *uuid.UUID           { return i.state.RevokedBy }
func (i *Invitation) RevokedAt() *time.Time           { return i.state.RevokedAt }
func (i *Invitation) RevokeReason() *string           { return i.state.RevokeReason }
func (i *Invitation) RoleIDs() []uuid.UUID            { return i.state.RoleIDs }
func (i *Invitation) Metadata() map[string]interface{} { return i.state.Metadata }
