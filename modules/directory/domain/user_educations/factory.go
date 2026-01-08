package user_educations

import (
	"time"

	"github.com/google/uuid"
)

type NewUserEducationParams struct {
	UserID      uuid.UUID
	School      string
	Degree      *string
	Major       *string
	FieldOfStudy *string
	StartDate   *time.Time
	EndDate     *time.Time
	IsCurrent   bool
	Description *string
	Grade       *string
	Activities  *string
	Achievements *string
}

func NewUserEducation(p NewUserEducationParams) (*UserEducation, error) {
	if err := validateUserEducationParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewUserEducationFromState(UserEducationState{
		ID:           id,
		UserID:       p.UserID,
		School:       p.School,
		Degree:       p.Degree,
		Major:        p.Major,
		FieldOfStudy: p.FieldOfStudy,
		StartDate:    p.StartDate,
		EndDate:      p.EndDate,
		IsCurrent:    p.IsCurrent,
		Description:  p.Description,
		Grade:        p.Grade,
		Activities:   p.Activities,
		Achievements: p.Achievements,
		CreatedAt:    now,
		UpdatedAt:    now,
	}), nil
}

func NewUserEducationFromState(st UserEducationState) *UserEducation {
	return &UserEducation{state: st}
}

func validateUserEducationParams(p NewUserEducationParams) error {
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if p.School == "" {
		return ErrSchoolRequired
	}
	return nil
}
