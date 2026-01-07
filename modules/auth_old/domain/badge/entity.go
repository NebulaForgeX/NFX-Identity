package badge

import (
	"time"

	"github.com/google/uuid"
)

type Badge struct {
	state BadgeState
}

type BadgeState struct {
	ID        uuid.UUID
	Editable  BadgeEditable
	DeletedAt *time.Time
}

type BadgeEditable struct {
	Name        string
	Description *string
	IconURL     *string
	Color       *string
	Category    *string // achievement, skill, community, special
	IsSystem    bool
}

func (b *Badge) ID() uuid.UUID              { return b.state.ID }
func (b *Badge) Editable() BadgeEditable    { return b.state.Editable }
func (b *Badge) DeletedAt() *time.Time       { return b.state.DeletedAt }

