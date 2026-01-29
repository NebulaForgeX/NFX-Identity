package action_requirements

import (
	"time"

	"github.com/google/uuid"
)

type NewActionRequirementParams struct {
	ActionID     uuid.UUID
	PermissionID uuid.UUID
	GroupID      int
}

func NewActionRequirement(p NewActionRequirementParams) (*ActionRequirement, error) {
	if err := validateActionRequirementParams(p); err != nil {
		return nil, err
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	groupID := p.GroupID
	if groupID <= 0 {
		groupID = 1
	}
	return NewActionRequirementFromState(ActionRequirementState{
		ID:           id,
		ActionID:     p.ActionID,
		PermissionID: p.PermissionID,
		GroupID:      groupID,
		CreatedAt:    now,
	}), nil
}

func NewActionRequirementFromState(st ActionRequirementState) *ActionRequirement {
	return &ActionRequirement{state: st}
}

func validateActionRequirementParams(p NewActionRequirementParams) error {
	if p.ActionID == uuid.Nil {
		return ErrActionIDRequired
	}
	if p.PermissionID == uuid.Nil {
		return ErrPermissionIDRequired
	}
	return nil
}
