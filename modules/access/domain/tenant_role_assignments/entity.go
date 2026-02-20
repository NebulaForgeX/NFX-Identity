package tenant_role_assignments

import (
	"time"

	"github.com/google/uuid"
)

type TenantRoleAssignment struct {
	state TenantRoleAssignmentState
}

type TenantRoleAssignmentState struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	TenantID     uuid.UUID
	TenantRoleID uuid.UUID
	AssignedAt   time.Time
	AssignedBy   *uuid.UUID
}

func (a *TenantRoleAssignment) ID() uuid.UUID             { return a.state.ID }
func (a *TenantRoleAssignment) UserID() uuid.UUID        { return a.state.UserID }
func (a *TenantRoleAssignment) TenantID() uuid.UUID       { return a.state.TenantID }
func (a *TenantRoleAssignment) TenantRoleID() uuid.UUID   { return a.state.TenantRoleID }
func (a *TenantRoleAssignment) AssignedAt() time.Time     { return a.state.AssignedAt }
func (a *TenantRoleAssignment) AssignedBy() *uuid.UUID     { return a.state.AssignedBy }
