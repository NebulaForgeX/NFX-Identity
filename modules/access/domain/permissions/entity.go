package permissions

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	state PermissionState
}

type PermissionState struct {
	ID          uuid.UUID
	Key         string
	Name        string
	Description *string
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (p *Permission) ID() uuid.UUID         { return p.state.ID }
func (p *Permission) Key() string           { return p.state.Key }
func (p *Permission) Name() string          { return p.state.Name }
func (p *Permission) Description() *string  { return p.state.Description }
func (p *Permission) IsSystem() bool        { return p.state.IsSystem }
func (p *Permission) CreatedAt() time.Time  { return p.state.CreatedAt }
func (p *Permission) UpdatedAt() time.Time  { return p.state.UpdatedAt }
func (p *Permission) DeletedAt() *time.Time { return p.state.DeletedAt }
