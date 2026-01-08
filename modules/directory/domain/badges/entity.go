package badges

import (
	"time"

	"github.com/google/uuid"
)

type Badge struct {
	state BadgeState
}

type BadgeState struct {
	ID          uuid.UUID
	Name        string
	Description *string
	IconURL     *string
	Color       *string
	Category    *string
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (b *Badge) ID() uuid.UUID            { return b.state.ID }
func (b *Badge) Name() string             { return b.state.Name }
func (b *Badge) Description() *string     { return b.state.Description }
func (b *Badge) IconURL() *string         { return b.state.IconURL }
func (b *Badge) Color() *string           { return b.state.Color }
func (b *Badge) Category() *string        { return b.state.Category }
func (b *Badge) IsSystem() bool           { return b.state.IsSystem }
func (b *Badge) CreatedAt() time.Time     { return b.state.CreatedAt }
func (b *Badge) UpdatedAt() time.Time     { return b.state.UpdatedAt }
func (b *Badge) DeletedAt() *time.Time    { return b.state.DeletedAt }
