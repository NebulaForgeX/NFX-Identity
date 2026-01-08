package member_roles

import (
	"time"

	"github.com/google/uuid"
)

type MemberRole struct {
	state MemberRoleState
}

type MemberRoleState struct {
	ID         uuid.UUID
	TenantID   uuid.UUID
	MemberID   uuid.UUID
	RoleID     uuid.UUID
	AssignedAt time.Time
	AssignedBy *uuid.UUID
	ExpiresAt  *time.Time
	Scope      *string
	RevokedAt  *time.Time
	RevokedBy  *uuid.UUID
	RevokeReason *string
}

func (mr *MemberRole) ID() uuid.UUID            { return mr.state.ID }
func (mr *MemberRole) TenantID() uuid.UUID       { return mr.state.TenantID }
func (mr *MemberRole) MemberID() uuid.UUID       { return mr.state.MemberID }
func (mr *MemberRole) RoleID() uuid.UUID         { return mr.state.RoleID }
func (mr *MemberRole) AssignedAt() time.Time     { return mr.state.AssignedAt }
func (mr *MemberRole) AssignedBy() *uuid.UUID    { return mr.state.AssignedBy }
func (mr *MemberRole) ExpiresAt() *time.Time     { return mr.state.ExpiresAt }
func (mr *MemberRole) Scope() *string            { return mr.state.Scope }
func (mr *MemberRole) RevokedAt() *time.Time     { return mr.state.RevokedAt }
func (mr *MemberRole) RevokedBy() *uuid.UUID     { return mr.state.RevokedBy }
func (mr *MemberRole) RevokeReason() *string     { return mr.state.RevokeReason }
