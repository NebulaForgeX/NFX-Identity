package results

import (
	"time"

	"nfxid/modules/directory/domain/user_occupations"

	"github.com/google/uuid"
)

type UserOccupationRO struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	Company         string
	Position        string
	Department      *string
	Industry        *string
	Location        *string
	EmploymentType  *string
	StartDate       *time.Time
	EndDate         *time.Time
	IsCurrent       bool
	Description     *string
	Responsibilities *string
	Achievements    *string
	SkillsUsed      []string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

// UserOccupationMapper 将 Domain UserOccupation 转换为 Application UserOccupationRO
func UserOccupationMapper(uo *user_occupations.UserOccupation) UserOccupationRO {
	if uo == nil {
		return UserOccupationRO{}
	}

	return UserOccupationRO{
		ID:              uo.ID(),
		UserID:          uo.UserID(),
		Company:         uo.Company(),
		Position:        uo.Position(),
		Department:      uo.Department(),
		Industry:        uo.Industry(),
		Location:        uo.Location(),
		EmploymentType:  uo.EmploymentType(),
		StartDate:       uo.StartDate(),
		EndDate:         uo.EndDate(),
		IsCurrent:       uo.IsCurrent(),
		Description:     uo.Description(),
		Responsibilities: uo.Responsibilities(),
		Achievements:    uo.Achievements(),
		SkillsUsed:      uo.SkillsUsed(),
		CreatedAt:       uo.CreatedAt(),
		UpdatedAt:       uo.UpdatedAt(),
		DeletedAt:       uo.DeletedAt(),
	}
}
