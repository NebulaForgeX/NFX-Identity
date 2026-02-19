package reqdto

import (
	arAppCommands "nfxid/modules/access/application/action_requirements/commands"

	"github.com/google/uuid"
)

type ActionRequirementByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type ActionRequirementByPermissionIDRequestDTO struct {
	PermissionID uuid.UUID `uri:"permission_id" validate:"required,uuid"`
}

type ActionRequirementCreateRequestDTO struct {
	ActionID     string `json:"action_id" validate:"required,uuid"`
	PermissionID string `json:"permission_id" validate:"required,uuid"`
	GroupID      int32  `json:"group_id,omitempty"`
}

func (r *ActionRequirementCreateRequestDTO) ToCreateCmd() (arAppCommands.CreateActionRequirementCmd, error) {
	actionID, err := uuid.Parse(r.ActionID)
	if err != nil {
		return arAppCommands.CreateActionRequirementCmd{}, err
	}
	permissionID, err := uuid.Parse(r.PermissionID)
	if err != nil {
		return arAppCommands.CreateActionRequirementCmd{}, err
	}
	return arAppCommands.CreateActionRequirementCmd{
		ActionID:     actionID,
		PermissionID: permissionID,
		GroupID:      int(r.GroupID),
	}, nil
}

type ActionRequirementDeleteRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}
