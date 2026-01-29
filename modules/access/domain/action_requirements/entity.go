package action_requirements

import (
	"time"

	"github.com/google/uuid"
)

type ActionRequirement struct {
	state ActionRequirementState
}

type ActionRequirementState struct {
	ID           uuid.UUID
	ActionID     uuid.UUID
	PermissionID uuid.UUID
	GroupID      int
	CreatedAt    time.Time
}

func (ar *ActionRequirement) ID() uuid.UUID           { return ar.state.ID }
func (ar *ActionRequirement) ActionID() uuid.UUID     { return ar.state.ActionID }
func (ar *ActionRequirement) PermissionID() uuid.UUID { return ar.state.PermissionID }
func (ar *ActionRequirement) GroupID() int            { return ar.state.GroupID }
func (ar *ActionRequirement) CreatedAt() time.Time    { return ar.state.CreatedAt }
