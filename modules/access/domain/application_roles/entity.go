package application_roles

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationRole struct {
	state ApplicationRoleState
}

type ApplicationRoleState struct {
	ID            uuid.UUID
	ApplicationID uuid.UUID
	RoleKey       string
	Name          *string
	CreatedAt     time.Time
}

func (r *ApplicationRole) ID() uuid.UUID            { return r.state.ID }
func (r *ApplicationRole) ApplicationID() uuid.UUID { return r.state.ApplicationID }
func (r *ApplicationRole) RoleKey() string          { return r.state.RoleKey }
func (r *ApplicationRole) Name() *string            { return r.state.Name }
func (r *ApplicationRole) CreatedAt() time.Time     { return r.state.CreatedAt }
