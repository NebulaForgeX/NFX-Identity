package views

import (
	"time"

	"github.com/google/uuid"
)

type OccupationView struct {
	ID               uuid.UUID
	ProfileID        uuid.UUID
	Company          string
	Position         string
	Department       *string
	Industry         *string
	Location         *string
	EmploymentType   *string
	StartDate        *time.Time
	EndDate          *time.Time
	IsCurrent        bool
	Description      *string
	Responsibilities *string
	Achievements     *string
	SkillsUsed       *string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}

