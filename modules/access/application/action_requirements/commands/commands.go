package commands

import (
	"github.com/google/uuid"
)

type CreateActionRequirementCmd struct {
	ActionID     uuid.UUID
	PermissionID uuid.UUID
	GroupID      int
}

type DeleteActionRequirementCmd struct {
	ActionRequirementID uuid.UUID
}

type DeleteByActionIDAndPermissionIDCmd struct {
	ActionID     uuid.UUID
	PermissionID uuid.UUID
}

type DeleteByActionIDCmd struct {
	ActionID uuid.UUID
}
