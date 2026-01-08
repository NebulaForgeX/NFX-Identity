package member_app_roles

import (
	"time"

	"github.com/google/uuid"
)

type MemberAppRole struct {
	state MemberAppRoleState
}

type MemberAppRoleState struct {
	ID           uuid.UUID
	MemberID     uuid.UUID
	AppID        uuid.UUID
	RoleID       uuid.UUID
	AssignedAt   time.Time
	AssignedBy   *uuid.UUID
	ExpiresAt    *time.Time
	RevokedAt    *time.Time
	RevokedBy    *uuid.UUID
	RevokeReason *string
}

func (mar *MemberAppRole) ID() uuid.UUID            { return mar.state.ID }
func (mar *MemberAppRole) MemberID() uuid.UUID       { return mar.state.MemberID }
func (mar *MemberAppRole) AppID() uuid.UUID          { return mar.state.AppID }
func (mar *MemberAppRole) RoleID() uuid.UUID         { return mar.state.RoleID }
func (mar *MemberAppRole) AssignedAt() time.Time     { return mar.state.AssignedAt }
func (mar *MemberAppRole) AssignedBy() *uuid.UUID    { return mar.state.AssignedBy }
func (mar *MemberAppRole) ExpiresAt() *time.Time     { return mar.state.ExpiresAt }
func (mar *MemberAppRole) RevokedAt() *time.Time     { return mar.state.RevokedAt }
func (mar *MemberAppRole) RevokedBy() *uuid.UUID     { return mar.state.RevokedBy }
func (mar *MemberAppRole) RevokeReason() *string     { return mar.state.RevokeReason }
