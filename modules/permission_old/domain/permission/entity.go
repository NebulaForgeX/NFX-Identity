package permission

import (
	"nfxid/enums"
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	state PermissionState
}

type PermissionState struct {
	ID        uuid.UUID
	Editable  PermissionEditable
	IsSystem  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type PermissionEditable struct {
	Tag         string
	Name        string
	Description string
	Category    enums.PermissionCategory
}

func (p *Permission) ID() uuid.UUID                { return p.state.ID }
func (p *Permission) Editable() PermissionEditable { return p.state.Editable }
func (p *Permission) IsSystem() bool               { return p.state.IsSystem }
func (p *Permission) CreatedAt() time.Time         { return p.state.CreatedAt }
func (p *Permission) UpdatedAt() time.Time         { return p.state.UpdatedAt }
func (p *Permission) DeletedAt() *time.Time        { return p.state.DeletedAt }

func (p *Permission) IsActive() bool {
	return p.DeletedAt() == nil
}
