package role

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	state RoleState
}

type RoleState struct {
	ID          uuid.UUID
	Editable    RoleEditable
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type RoleEditable struct {
	Name        string
	Description *string
	Permissions []string
}

func (r *Role) ID() uuid.UUID              { return r.state.ID }
func (r *Role) Editable() RoleEditable      { return r.state.Editable }
func (r *Role) IsSystem() bool              { return r.state.IsSystem }
func (r *Role) CreatedAt() time.Time        { return r.state.CreatedAt }
func (r *Role) UpdatedAt() time.Time        { return r.state.UpdatedAt }
func (r *Role) DeletedAt() *time.Time       { return r.state.DeletedAt }

func (r *Role) HasPermission(permission string) bool {
	for _, p := range r.Editable().Permissions {
		if p == permission {
			return true
		}
	}
	return false
}
