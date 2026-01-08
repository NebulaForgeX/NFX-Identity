package user_occupations

import (
	"time"

	"github.com/google/uuid"
)

type NewUserOccupationParams struct {
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
}

func NewUserOccupation(p NewUserOccupationParams) (*UserOccupation, error) {
	if err := validateUserOccupationParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewUserOccupationFromState(UserOccupationState{
		ID:              id,
		UserID:          p.UserID,
		Company:         p.Company,
		Position:        p.Position,
		Department:      p.Department,
		Industry:        p.Industry,
		Location:        p.Location,
		EmploymentType:  p.EmploymentType,
		StartDate:       p.StartDate,
		EndDate:         p.EndDate,
		IsCurrent:       p.IsCurrent,
		Description:     p.Description,
		Responsibilities: p.Responsibilities,
		Achievements:    p.Achievements,
		SkillsUsed:      p.SkillsUsed,
		CreatedAt:       now,
		UpdatedAt:       now,
	}), nil
}

func NewUserOccupationFromState(st UserOccupationState) *UserOccupation {
	return &UserOccupation{state: st}
}

func validateUserOccupationParams(p NewUserOccupationParams) error {
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if p.Company == "" {
		return ErrCompanyRequired
	}
	if p.Position == "" {
		return ErrPositionRequired
	}
	return nil
}
