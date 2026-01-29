package mapper

import (
	"time"

	"nfxid/modules/access/domain/action_requirements"

	"github.com/google/uuid"
)

// ActionRequirementModel matches access.action_requirements table (includes group_id; dbgen may not)
type ActionRequirementModel struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	ActionID     uuid.UUID `gorm:"type:uuid;not null"`
	PermissionID uuid.UUID `gorm:"type:uuid;not null"`
	GroupID      int       `gorm:"type:integer;not null;default:1"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

func (ActionRequirementModel) TableName() string { return "access.action_requirements" }

func ActionRequirementDomainToModel(ar *action_requirements.ActionRequirement) *ActionRequirementModel {
	if ar == nil {
		return nil
	}
	return &ActionRequirementModel{
		ID:           ar.ID(),
		ActionID:     ar.ActionID(),
		PermissionID: ar.PermissionID(),
		GroupID:      ar.GroupID(),
		CreatedAt:    ar.CreatedAt(),
	}
}

func ActionRequirementModelToDomain(m *ActionRequirementModel) *action_requirements.ActionRequirement {
	if m == nil {
		return nil
	}
	state := action_requirements.ActionRequirementState{
		ID:           m.ID,
		ActionID:     m.ActionID,
		PermissionID: m.PermissionID,
		GroupID:      m.GroupID,
		CreatedAt:    m.CreatedAt,
	}
	return action_requirements.NewActionRequirementFromState(state)
}
