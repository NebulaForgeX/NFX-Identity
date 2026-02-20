package tenant_roles

import (
	"time"

	"github.com/google/uuid"
)

type TenantRole struct {
	state TenantRoleState
}

type TenantRoleState struct {
	ID        uuid.UUID
	TenantID  uuid.UUID
	RoleKey   string
	Name      *string
	CreatedAt time.Time
}

func (r *TenantRole) ID() uuid.UUID        { return r.state.ID }
func (r *TenantRole) TenantID() uuid.UUID  { return r.state.TenantID }
func (r *TenantRole) RoleKey() string      { return r.state.RoleKey }
func (r *TenantRole) Name() *string        { return r.state.Name }
func (r *TenantRole) CreatedAt() time.Time { return r.state.CreatedAt }
