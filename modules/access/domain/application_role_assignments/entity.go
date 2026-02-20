package application_role_assignments

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationRoleAssignment struct {
	state ApplicationRoleAssignmentState
}

type ApplicationRoleAssignmentState struct {
	ID                uuid.UUID
	UserID            uuid.UUID
	ApplicationID     uuid.UUID
	ApplicationRoleID uuid.UUID
	AssignedAt        time.Time
	AssignedBy        *uuid.UUID
}

func (a *ApplicationRoleAssignment) ID() uuid.UUID              { return a.state.ID }
func (a *ApplicationRoleAssignment) UserID() uuid.UUID          { return a.state.UserID }
func (a *ApplicationRoleAssignment) ApplicationID() uuid.UUID    { return a.state.ApplicationID }
func (a *ApplicationRoleAssignment) ApplicationRoleID() uuid.UUID { return a.state.ApplicationRoleID }
func (a *ApplicationRoleAssignment) AssignedAt() time.Time      { return a.state.AssignedAt }
func (a *ApplicationRoleAssignment) AssignedBy() *uuid.UUID     { return a.state.AssignedBy }
