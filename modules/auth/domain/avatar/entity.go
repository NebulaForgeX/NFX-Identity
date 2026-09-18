package avatar

import (
	"time"

	"github.com/google/uuid"
)

type Avatar struct{ state State }

type State struct {
	ID        uuid.UUID
	ProfileID uuid.UUID
	ImageID   uuid.UUID
	IsActive  bool
	Kind      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewFromState(st State) *Avatar     { return &Avatar{state: st} }
func (a *Avatar) ID() uuid.UUID         { return a.state.ID }
func (a *Avatar) ProfileID() uuid.UUID  { return a.state.ProfileID }
func (a *Avatar) ImageID() uuid.UUID    { return a.state.ImageID }
func (a *Avatar) IsActive() bool        { return a.state.IsActive }
func (a *Avatar) Kind() string          { return a.state.Kind }
func (a *Avatar) CreatedAt() time.Time  { return a.state.CreatedAt }
func (a *Avatar) UpdatedAt() time.Time  { return a.state.UpdatedAt }
func (a *Avatar) DeletedAt() *time.Time { return a.state.DeletedAt }
func (a *Avatar) State() State          { return a.state }
