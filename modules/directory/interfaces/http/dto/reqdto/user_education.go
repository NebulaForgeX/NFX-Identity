package reqdto

import (
	userEducationAppCommands "nfxid/modules/directory/application/user_educations/commands"

	"github.com/google/uuid"
)

type UserEducationCreateRequestDTO struct {
	UserID       uuid.UUID `json:"user_id" validate:"required,uuid"`
	School       string    `json:"school" validate:"required"`
	Degree       *string   `json:"degree,omitempty"`
	Major        *string   `json:"major,omitempty"`
	FieldOfStudy *string   `json:"field_of_study,omitempty"`
	StartDate    *string   `json:"start_date,omitempty"`
	EndDate      *string   `json:"end_date,omitempty"`
	IsCurrent    bool      `json:"is_current"`
	Description  *string   `json:"description,omitempty"`
	Grade        *string   `json:"grade,omitempty"`
	Activities   *string   `json:"activities,omitempty"`
	Achievements *string   `json:"achievements,omitempty"`
}

type UserEducationByIDRequestDTO struct {
	UserEducationID uuid.UUID `uri:"user_education_id" validate:"required,uuid"`
}

type UserEducationUpdateRequestDTO struct {
	UserEducationID uuid.UUID `uri:"user_education_id" validate:"required,uuid"`
	School       string    `json:"school" validate:"required"`
	Degree       *string   `json:"degree,omitempty"`
	Major        *string   `json:"major,omitempty"`
	FieldOfStudy *string   `json:"field_of_study,omitempty"`
	StartDate    *string   `json:"start_date,omitempty"`
	EndDate      *string   `json:"end_date,omitempty"`
	IsCurrent    bool      `json:"is_current"`
	Description  *string   `json:"description,omitempty"`
	Grade        *string   `json:"grade,omitempty"`
	Activities   *string   `json:"activities,omitempty"`
	Achievements *string   `json:"achievements,omitempty"`
}

func (r *UserEducationCreateRequestDTO) ToCreateCmd() userEducationAppCommands.CreateUserEducationCmd {
	return userEducationAppCommands.CreateUserEducationCmd{
		UserID:       r.UserID,
		School:       r.School,
		Degree:       r.Degree,
		Major:        r.Major,
		FieldOfStudy: r.FieldOfStudy,
		StartDate:    r.StartDate,
		EndDate:      r.EndDate,
		IsCurrent:    r.IsCurrent,
		Description:  r.Description,
		Grade:        r.Grade,
		Activities:   r.Activities,
		Achievements: r.Achievements,
	}
}

func (r *UserEducationUpdateRequestDTO) ToUpdateCmd() userEducationAppCommands.UpdateUserEducationCmd {
	return userEducationAppCommands.UpdateUserEducationCmd{
		UserEducationID: r.UserEducationID,
		School:          r.School,
		Degree:          r.Degree,
		Major:           r.Major,
		FieldOfStudy:    r.FieldOfStudy,
		StartDate:       r.StartDate,
		EndDate:         r.EndDate,
		IsCurrent:       r.IsCurrent,
		Description:     r.Description,
		Grade:           r.Grade,
		Activities:      r.Activities,
		Achievements:    r.Achievements,
	}
}
