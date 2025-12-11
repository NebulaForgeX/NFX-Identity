package views

import (
	"time"

	educationDomainViews "nfxid/modules/auth/domain/profile_education/views"

	"github.com/google/uuid"
)

type EducationView struct {
	ID           uuid.UUID  `json:"id"`
	ProfileID    uuid.UUID  `json:"profile_id"`
	School       string     `json:"school"`
	Degree       *string    `json:"degree"`
	Major        *string    `json:"major"`
	FieldOfStudy *string    `json:"field_of_study"`
	StartDate    *time.Time `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	IsCurrent    bool       `json:"is_current"`
	Description  *string    `json:"description"`
	Grade        *string    `json:"grade"`
	Activities   *string    `json:"activities"`
	Achievements *string    `json:"achievements"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

// EducationViewMapper 将 Domain EducationView 转换为 Application EducationView
func EducationViewMapper(v educationDomainViews.EducationView) EducationView {
	return EducationView{
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
