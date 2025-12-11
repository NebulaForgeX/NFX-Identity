package education

import (
	"time"

	"github.com/google/uuid"
)

type Education struct {
	state EducationState
}

type EducationState struct {
	ID        uuid.UUID
	ProfileID uuid.UUID
	Editable  EducationEditable
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type EducationEditable struct {
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
}

func (e *Education) ID() uuid.UUID              { return e.state.ID }
func (e *Education) ProfileID() uuid.UUID         { return e.state.ProfileID }
func (e *Education) Editable() EducationEditable  { return e.state.Editable }
func (e *Education) CreatedAt() time.Time         { return e.state.CreatedAt }
func (e *Education) UpdatedAt() time.Time         { return e.state.UpdatedAt }
func (e *Education) DeletedAt() *time.Time        { return e.state.DeletedAt }
