package user_occupations

import (
	"time"

	"github.com/google/uuid"
)

type UserOccupation struct {
	state UserOccupationState
}

type UserOccupationState struct {
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

func (uo *UserOccupation) ID() uuid.UUID              { return uo.state.ID }
func (uo *UserOccupation) UserID() uuid.UUID          { return uo.state.UserID }
func (uo *UserOccupation) Company() string            { return uo.state.Company }
func (uo *UserOccupation) Position() string           { return uo.state.Position }
func (uo *UserOccupation) Department() *string        { return uo.state.Department }
func (uo *UserOccupation) Industry() *string          { return uo.state.Industry }
func (uo *UserOccupation) Location() *string          { return uo.state.Location }
func (uo *UserOccupation) EmploymentType() *string    { return uo.state.EmploymentType }
func (uo *UserOccupation) StartDate() *time.Time      { return uo.state.StartDate }
func (uo *UserOccupation) EndDate() *time.Time        { return uo.state.EndDate }
func (uo *UserOccupation) IsCurrent() bool            { return uo.state.IsCurrent }
func (uo *UserOccupation) Description() *string       { return uo.state.Description }
func (uo *UserOccupation) Responsibilities() *string  { return uo.state.Responsibilities }
func (uo *UserOccupation) Achievements() *string      { return uo.state.Achievements }
func (uo *UserOccupation) SkillsUsed() []string       { return uo.state.SkillsUsed }
func (uo *UserOccupation) CreatedAt() time.Time       { return uo.state.CreatedAt }
func (uo *UserOccupation) UpdatedAt() time.Time       { return uo.state.UpdatedAt }
func (uo *UserOccupation) DeletedAt() *time.Time      { return uo.state.DeletedAt }
