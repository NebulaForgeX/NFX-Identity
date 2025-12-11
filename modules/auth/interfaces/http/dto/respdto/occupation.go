package respdto

import (
	"time"

	occupationAppViews "nfxid/modules/auth/application/occupation/views"

	"github.com/google/uuid"
)

type OccupationDTO struct {
	ID               uuid.UUID  `json:"id"`
	ProfileID        uuid.UUID  `json:"profile_id"`
	Company          string     `json:"company"`
	Position         string     `json:"position"`
	Department       *string    `json:"department,omitempty"`
	Industry         *string    `json:"industry,omitempty"`
	Location         *string    `json:"location,omitempty"`
	EmploymentType   *string    `json:"employment_type,omitempty"`
	StartDate        *time.Time `json:"start_date,omitempty"`
	EndDate          *time.Time `json:"end_date,omitempty"`
	IsCurrent        bool       `json:"is_current"`
	Description      *string    `json:"description,omitempty"`
	Responsibilities *string    `json:"responsibilities,omitempty"`
	Achievements     *string    `json:"achievements,omitempty"`
	SkillsUsed       *string    `json:"skills_used,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

// OccupationViewToDTO converts application OccupationView to response DTO
func OccupationViewToDTO(v *occupationAppViews.OccupationView) *OccupationDTO {
	if v == nil {
		return nil
	}

	return &OccupationDTO{
		ID:               v.ID,
		ProfileID:        v.ProfileID,
		Company:          v.Company,
		Position:         v.Position,
		Department:       v.Department,
		Industry:         v.Industry,
		Location:         v.Location,
		EmploymentType:   v.EmploymentType,
		StartDate:        v.StartDate,
		EndDate:          v.EndDate,
		IsCurrent:        v.IsCurrent,
		Description:      v.Description,
		Responsibilities: v.Responsibilities,
		Achievements:     v.Achievements,
		SkillsUsed:       v.SkillsUsed,
		CreatedAt:        v.CreatedAt,
		UpdatedAt:        v.UpdatedAt,
		DeletedAt:        v.DeletedAt,
	}
}

// OccupationListViewToDTO converts list of OccupationView to DTOs
func OccupationListViewToDTO(views []occupationAppViews.OccupationView) []OccupationDTO {
	dtos := make([]OccupationDTO, len(views))
	for i, v := range views {
		if dto := OccupationViewToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
