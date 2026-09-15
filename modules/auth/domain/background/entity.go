package background

import (
	"time"

	"github.com/google/uuid"
)

type Background struct{ state State }

type State struct {
	ID        uuid.UUID
	ProfileID uuid.UUID
	ImageID   uuid.UUID
	SortOrder int
	Kind      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewFromState(st State) *Background { return &Background{state: st} }
func (b *Background) ID() uuid.UUID        { return b.state.ID }
func (b *Background) ProfileID() uuid.UUID { return b.state.ProfileID }
func (b *Background) ImageID() uuid.UUID   { return b.state.ImageID }
func (b *Background) SortOrder() int       { return b.state.SortOrder }
func (b *Background) Kind() string         { return b.state.Kind }
func (b *Background) CreatedAt() time.Time { return b.state.CreatedAt }
func (b *Background) UpdatedAt() time.Time { return b.state.UpdatedAt }
func (b *Background) State() State         { return b.state }
