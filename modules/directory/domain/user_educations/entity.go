package user_educations

import (
	"time"

	"github.com/google/uuid"
)

type UserEducation struct {
	state UserEducationState
}

type UserEducationState struct {
	ID          uuid.UUID
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
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (ue *UserEducation) ID() uuid.UUID            { return ue.state.ID }
func (ue *UserEducation) UserID() uuid.UUID        { return ue.state.UserID }
func (ue *UserEducation) School() string           { return ue.state.School }
func (ue *UserEducation) Degree() *string          { return ue.state.Degree }
func (ue *UserEducation) Major() *string           { return ue.state.Major }
func (ue *UserEducation) FieldOfStudy() *string    { return ue.state.FieldOfStudy }
func (ue *UserEducation) StartDate() *time.Time    { return ue.state.StartDate }
func (ue *UserEducation) EndDate() *time.Time      { return ue.state.EndDate }
func (ue *UserEducation) IsCurrent() bool          { return ue.state.IsCurrent }
func (ue *UserEducation) Description() *string     { return ue.state.Description }
func (ue *UserEducation) Grade() *string           { return ue.state.Grade }
func (ue *UserEducation) Activities() *string      { return ue.state.Activities }
func (ue *UserEducation) Achievements() *string    { return ue.state.Achievements }
func (ue *UserEducation) CreatedAt() time.Time     { return ue.state.CreatedAt }
func (ue *UserEducation) UpdatedAt() time.Time     { return ue.state.UpdatedAt }
func (ue *UserEducation) DeletedAt() *time.Time    { return ue.state.DeletedAt }
