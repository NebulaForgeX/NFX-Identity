package respdto

import (
	"time"

	userEducationAppResult "nfxid/modules/directory/application/user_educations/results"

	"github.com/google/uuid"
)

type UserEducationDTO struct {
	ID          uuid.UUID  `json:"id"`
	UserID      uuid.UUID  `json:"user_id"`
	School      string     `json:"school"`
	Degree      *string    `json:"degree,omitempty"`
	Major       *string    `json:"major,omitempty"`
	FieldOfStudy *string    `json:"field_of_study,omitempty"`
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	IsCurrent   bool       `json:"is_current"`
	Description *string    `json:"description,omitempty"`
	Grade       *string    `json:"grade,omitempty"`
	Activities  *string    `json:"activities,omitempty"`
	Achievements *string    `json:"achievements,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// UserEducationROToDTO converts application UserEducationRO to response DTO
func UserEducationROToDTO(v *userEducationAppResult.UserEducationRO) *UserEducationDTO {
	if v == nil {
		return nil
	}

	return &UserEducationDTO{
		ID:          v.ID,
		UserID:      v.UserID,
		School:      v.School,
		Degree:      v.Degree,
		Major:       v.Major,
		FieldOfStudy: v.FieldOfStudy,
		StartDate:   v.StartDate,
		EndDate:     v.EndDate,
		IsCurrent:   v.IsCurrent,
		Description: v.Description,
		Grade:       v.Grade,
		Activities:  v.Activities,
		Achievements: v.Achievements,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		DeletedAt:   v.DeletedAt,
	}
}

// UserEducationListROToDTO converts list of UserEducationRO to DTOs
func UserEducationListROToDTO(results []userEducationAppResult.UserEducationRO) []UserEducationDTO {
	dtos := make([]UserEducationDTO, len(results))
	for i, v := range results {
		if dto := UserEducationROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
