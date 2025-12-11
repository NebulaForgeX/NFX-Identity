package occupation

import (
	"time"

	"github.com/google/uuid"
)

type Occupation struct {
	state OccupationState
}

type OccupationState struct {
	ID        uuid.UUID
	ProfileID uuid.UUID
	Editable  OccupationEditable
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type OccupationEditable struct {
	Company          string
	Position         string
	Department       *string
	Industry         *string
	Location         *string
	EmploymentType   *string // full-time, part-time, contract, etc.
	StartDate        *time.Time
	EndDate          *time.Time
	IsCurrent        bool
	Description      *string
	Responsibilities *string
	Achievements     *string
	SkillsUsed       []string // PostgreSQL text[]
}

func (o *Occupation) ID() uuid.UUID                { return o.state.ID }
func (o *Occupation) ProfileID() uuid.UUID          { return o.state.ProfileID }
func (o *Occupation) Editable() OccupationEditable  { return o.state.Editable }
func (o *Occupation) CreatedAt() time.Time          { return o.state.CreatedAt }
func (o *Occupation) UpdatedAt() time.Time          { return o.state.UpdatedAt }
func (o *Occupation) DeletedAt() *time.Time         { return o.state.DeletedAt }
