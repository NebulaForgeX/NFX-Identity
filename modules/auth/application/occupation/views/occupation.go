package views

import (
	"time"

	occupationDomainViews "nebulaid/modules/auth/domain/occupation/views"

	"github.com/google/uuid"
)

type OccupationView struct {
	ID               uuid.UUID  `json:"id"`
	ProfileID        uuid.UUID  `json:"profile_id"`
	Company          string     `json:"company"`
	Position         string     `json:"position"`
	Department       *string    `json:"department"`
	Industry         *string    `json:"industry"`
	Location         *string    `json:"location"`
	EmploymentType   *string    `json:"employment_type"`
	StartDate        *time.Time `json:"start_date"`
	EndDate          *time.Time `json:"end_date"`
	IsCurrent        bool       `json:"is_current"`
	Description      *string    `json:"description"`
	Responsibilities *string    `json:"responsibilities"`
	Achievements     *string    `json:"achievements"`
	SkillsUsed       *string    `json:"skills_used"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

// OccupationViewMapper 将 Domain OccupationView 转换为 Application OccupationView
func OccupationViewMapper(v occupationDomainViews.OccupationView) OccupationView {
	return OccupationView{
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
