package reqdto

import (
	userOccupationAppCommands "nfxid/modules/directory/application/user_occupations/commands"

	"github.com/google/uuid"
)

type UserOccupationCreateRequestDTO struct {
	UserID          uuid.UUID `json:"user_id" validate:"required,uuid"`
	Company         string    `json:"company" validate:"required"`
	Position        string    `json:"position" validate:"required"`
	Department      *string   `json:"department,omitempty"`
	Industry        *string   `json:"industry,omitempty"`
	Location        *string   `json:"location,omitempty"`
	EmploymentType  *string   `json:"employment_type,omitempty"`
	StartDate       *string   `json:"start_date,omitempty"`
	EndDate         *string   `json:"end_date,omitempty"`
	IsCurrent       bool      `json:"is_current"`
	Description     *string   `json:"description,omitempty"`
	Responsibilities *string   `json:"responsibilities,omitempty"`
	Achievements    *string   `json:"achievements,omitempty"`
	SkillsUsed      []string  `json:"skills_used,omitempty"`
}

type UserOccupationUpdateRequestDTO struct {
	ID              uuid.UUID `params:"id" validate:"required,uuid"`
	Company         string    `json:"company" validate:"required"`
	Position        string    `json:"position" validate:"required"`
	Department      *string   `json:"department,omitempty"`
	Industry        *string   `json:"industry,omitempty"`
	Location        *string   `json:"location,omitempty"`
	EmploymentType  *string   `json:"employment_type,omitempty"`
	StartDate       *string   `json:"start_date,omitempty"`
	EndDate         *string   `json:"end_date,omitempty"`
	IsCurrent       bool      `json:"is_current"`
	Description     *string   `json:"description,omitempty"`
	Responsibilities *string   `json:"responsibilities,omitempty"`
	Achievements    *string   `json:"achievements,omitempty"`
	SkillsUsed      []string  `json:"skills_used,omitempty"`
}

func (r *UserOccupationCreateRequestDTO) ToCreateCmd() userOccupationAppCommands.CreateUserOccupationCmd {
	return userOccupationAppCommands.CreateUserOccupationCmd{
		UserID:          r.UserID,
		Company:         r.Company,
		Position:        r.Position,
		Department:      r.Department,
		Industry:        r.Industry,
		Location:        r.Location,
		EmploymentType:  r.EmploymentType,
		StartDate:       r.StartDate,
		EndDate:         r.EndDate,
		IsCurrent:       r.IsCurrent,
		Description:     r.Description,
		Responsibilities: r.Responsibilities,
		Achievements:    r.Achievements,
		SkillsUsed:      r.SkillsUsed,
	}
}

func (r *UserOccupationUpdateRequestDTO) ToUpdateCmd() userOccupationAppCommands.UpdateUserOccupationCmd {
	return userOccupationAppCommands.UpdateUserOccupationCmd{
		UserOccupationID: r.ID,
		Company:          r.Company,
		Position:         r.Position,
		Department:       r.Department,
		Industry:         r.Industry,
		Location:         r.Location,
		EmploymentType:   r.EmploymentType,
		StartDate:        r.StartDate,
		EndDate:          r.EndDate,
		IsCurrent:        r.IsCurrent,
		Description:      r.Description,
		Responsibilities: r.Responsibilities,
		Achievements:     r.Achievements,
		SkillsUsed:       r.SkillsUsed,
	}
}
