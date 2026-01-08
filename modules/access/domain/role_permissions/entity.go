package role_permissions

import (
	"time"

	"github.com/google/uuid"
)

type RolePermission struct {
	state RolePermissionState
}

type RolePermissionState struct {
	ID           uuid.UUID
	RoleID       uuid.UUID
	PermissionID uuid.UUID
	CreatedAt    time.Time
	CreatedBy    *uuid.UUID
}

func (rp *RolePermission) ID() uuid.UUID           { return rp.state.ID }
func (rp *RolePermission) RoleID() uuid.UUID       { return rp.state.RoleID }
func (rp *RolePermission) PermissionID() uuid.UUID { return rp.state.PermissionID }
func (rp *RolePermission) CreatedAt() time.Time    { return rp.state.CreatedAt }
func (rp *RolePermission) CreatedBy() *uuid.UUID   { return rp.state.CreatedBy }
