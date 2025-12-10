package respdto

import (
	"time"

	educationAppViews "nebulaid/modules/auth/application/education/views"

	"github.com/google/uuid"
)

type EducationDTO struct {
	ID           uuid.UUID  `json:"id"`
	ProfileID    uuid.UUID  `json:"profile_id"`
	School       string     `json:"school"`
	Degree       *string    `json:"degree,omitempty"`
	Major        *string    `json:"major,omitempty"`
	FieldOfStudy *string    `json:"field_of_study,omitempty"`
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	IsCurrent    bool       `json:"is_current"`
	Description  *string    `json:"description,omitempty"`
	Grade        *string    `json:"grade,omitempty"`
	Activities   *string    `json:"activities,omitempty"`
	Achievements *string    `json:"achievements,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

// EducationViewToDTO converts application EducationView to response DTO
func EducationViewToDTO(v *educationAppViews.EducationView) *EducationDTO {
	if v == nil {
		return nil
	}

	return &EducationDTO{
		ID:           v.ID,
		ProfileID:    v.ProfileID,
		School:       v.School,
		Degree:       v.Degree,
		Major:        v.Major,
		FieldOfStudy: v.FieldOfStudy,
		StartDate:    v.StartDate,
		EndDate:      v.EndDate,
		IsCurrent:    v.IsCurrent,
		Description:  v.Description,
		Grade:        v.Grade,
		Activities:   v.Activities,
		Achievements: v.Achievements,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
		DeletedAt:    v.DeletedAt,
	}
}

// EducationListViewToDTO converts list of EducationView to DTOs
func EducationListViewToDTO(views []educationAppViews.EducationView) []EducationDTO {
	dtos := make([]EducationDTO, len(views))
	for i, v := range views {
		if dto := EducationViewToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
