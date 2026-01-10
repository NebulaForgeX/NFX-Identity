package results

import (
	"time"

	"nfxid/modules/access/domain/scope_permissions"

	"github.com/google/uuid"
)

type ScopePermissionRO struct {
	ID           uuid.UUID
	Scope        string
	PermissionID uuid.UUID
	CreatedAt    time.Time
}

// ScopePermissionMapper 将 Domain ScopePermission 转换为 Application ScopePermissionRO
func ScopePermissionMapper(sp *scope_permissions.ScopePermission) ScopePermissionRO {
	if sp == nil {
		return ScopePermissionRO{}
	}

	return ScopePermissionRO{
		ID:           sp.ID(),
		Scope:        sp.Scope(),
		PermissionID: sp.PermissionID(),
		CreatedAt:    sp.CreatedAt(),
	}
}
