package scope_permissions

import (
	"time"

	"github.com/google/uuid"
)

type ScopePermission struct {
	state ScopePermissionState
}

type ScopePermissionState struct {
	ID           uuid.UUID
	Scope        string
	PermissionID uuid.UUID
	CreatedAt    time.Time
}

func (sp *ScopePermission) ID() uuid.UUID        { return sp.state.ID }
func (sp *ScopePermission) Scope() string        { return sp.state.Scope }
func (sp *ScopePermission) PermissionID() uuid.UUID { return sp.state.PermissionID }
func (sp *ScopePermission) CreatedAt() time.Time { return sp.state.CreatedAt }
