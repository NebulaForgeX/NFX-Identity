package respdto

import (
	"time"

	userOccupationAppResult "nfxid/modules/directory/application/user_occupations/results"

	"github.com/google/uuid"
)

type UserOccupationDTO struct {
	ID              uuid.UUID  `json:"id"`
	UserID          uuid.UUID  `json:"user_id"`
	Company         string     `json:"company"`
	Position        string     `json:"position"`
	Department      *string    `json:"department,omitempty"`
	Industry        *string    `json:"industry,omitempty"`
	Location        *string    `json:"location,omitempty"`
	EmploymentType  *string    `json:"employment_type,omitempty"`
	StartDate       *time.Time `json:"start_date,omitempty"`
	EndDate         *time.Time `json:"end_date,omitempty"`
	IsCurrent       bool       `json:"is_current"`
	Description     *string    `json:"description,omitempty"`
	Responsibilities *string    `json:"responsibilities,omitempty"`
	Achievements    *string    `json:"achievements,omitempty"`
	SkillsUsed      []string   `json:"skills_used,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

// UserOccupationROToDTO converts application UserOccupationRO to response DTO
func UserOccupationROToDTO(v *userOccupationAppResult.UserOccupationRO) *UserOccupationDTO {
	if v == nil {
		return nil
	}

	return &UserOccupationDTO{
		ID:              v.ID,
		UserID:          v.UserID,
		Company:         v.Company,
		Position:        v.Position,
		Department:      v.Department,
		Industry:        v.Industry,
		Location:        v.Location,
		EmploymentType:  v.EmploymentType,
		StartDate:       v.StartDate,
		EndDate:         v.EndDate,
		IsCurrent:       v.IsCurrent,
		Description:     v.Description,
		Responsibilities: v.Responsibilities,
		Achievements:    v.Achievements,
		SkillsUsed:      v.SkillsUsed,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
		DeletedAt:       v.DeletedAt,
	}
}

// UserOccupationListROToDTO converts list of UserOccupationRO to DTOs
func UserOccupationListROToDTO(results []userOccupationAppResult.UserOccupationRO) []UserOccupationDTO {
	dtos := make([]UserOccupationDTO, len(results))
	for i, v := range results {
		if dto := UserOccupationROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
