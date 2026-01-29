package respdto

import (
	"time"

	arAppResult "nfxid/modules/access/application/action_requirements/results"

	"github.com/google/uuid"
)

type ActionRequirementDTO struct {
	ID           uuid.UUID `json:"id"`
	ActionID     uuid.UUID `json:"action_id"`
	PermissionID uuid.UUID `json:"permission_id"`
	GroupID      int      `json:"group_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func ActionRequirementROToDTO(v *arAppResult.ActionRequirementRO) *ActionRequirementDTO {
	if v == nil {
		return nil
	}
	return &ActionRequirementDTO{
		ID:           v.ID,
		ActionID:     v.ActionID,
		PermissionID: v.PermissionID,
		GroupID:      v.GroupID,
		CreatedAt:    v.CreatedAt,
	}
}

func ActionRequirementListROToDTO(results []arAppResult.ActionRequirementRO) []ActionRequirementDTO {
	dtos := make([]ActionRequirementDTO, len(results))
	for i, v := range results {
		if dto := ActionRequirementROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
