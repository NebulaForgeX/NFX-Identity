package results

import (
	"time"

	"nfxid/modules/access/domain/action_requirements"

	"github.com/google/uuid"
)

type ActionRequirementRO struct {
	ID           uuid.UUID
	ActionID     uuid.UUID
	PermissionID uuid.UUID
	GroupID      int
	CreatedAt    time.Time
}

func ActionRequirementMapper(ar *action_requirements.ActionRequirement) ActionRequirementRO {
	if ar == nil {
		return ActionRequirementRO{}
	}
	return ActionRequirementRO{
		ID:           ar.ID(),
		ActionID:     ar.ActionID(),
		PermissionID: ar.PermissionID(),
		GroupID:      ar.GroupID(),
		CreatedAt:    ar.CreatedAt(),
	}
}
