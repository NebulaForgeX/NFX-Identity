package views

import (
	"time"

	"github.com/google/uuid"
)

type EducationView struct {
	ID           uuid.UUID
	ProfileID    uuid.UUID
	School       string
	Degree       *string
	Major        *string
	FieldOfStudy *string
	StartDate    *time.Time
	EndDate      *time.Time
	IsCurrent    bool
	Description  *string
	Grade        *string
	Activities   *string
	Achievements *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

