package reqdto

import (
	scopePermissionAppCommands "nfxid/modules/access/application/scope_permissions/commands"

	"github.com/google/uuid"
)

type ScopePermissionCreateRequestDTO struct {
	Scope        string     `json:"scope" validate:"required"`
	PermissionID uuid.UUID  `json:"permission_id" validate:"required"`
}

type ScopePermissionByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type ScopePermissionByScopeAndPermissionRequestDTO struct {
	Scope        string    `json:"scope" validate:"required"`
	PermissionID uuid.UUID `json:"permission_id" validate:"required"`
}

func (r *ScopePermissionCreateRequestDTO) ToCreateCmd() scopePermissionAppCommands.CreateScopePermissionCmd {
	return scopePermissionAppCommands.CreateScopePermissionCmd{
		Scope:        r.Scope,
		PermissionID: r.PermissionID,
	}
}
