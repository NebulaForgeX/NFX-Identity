package roles

import (
	"time"

	"github.com/google/uuid"
)

type ScopeType string

const (
	ScopeTypeTenant ScopeType = "TENANT"
	ScopeTypeApp    ScopeType = "APP"
	ScopeTypeGlobal ScopeType = "GLOBAL"
)

type Role struct {
	state RoleState
}

type RoleState struct {
	ID          uuid.UUID
	Key         string
	Name        string
	Description *string
	ScopeType   ScopeType
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (r *Role) ID() uuid.UUID        { return r.state.ID }
func (r *Role) Key() string          { return r.state.Key }
func (r *Role) Name() string         { return r.state.Name }
func (r *Role) Description() *string { return r.state.Description }
func (r *Role) ScopeType() ScopeType { return r.state.ScopeType }
func (r *Role) IsSystem() bool       { return r.state.IsSystem }
func (r *Role) CreatedAt() time.Time { return r.state.CreatedAt }
func (r *Role) UpdatedAt() time.Time { return r.state.UpdatedAt }
func (r *Role) DeletedAt() *time.Time { return r.state.DeletedAt }
