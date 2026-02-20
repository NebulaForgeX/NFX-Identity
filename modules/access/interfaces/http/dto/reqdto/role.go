package reqdto

import (
	roleAppCommands "nfxid/modules/access/application/roles/commands"
	roleDomain "nfxid/modules/access/domain/roles"

	"github.com/google/uuid"
)

type RoleCreateRequestDTO struct {
	Key         string  `json:"key" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ScopeType   string  `json:"scope_type,omitempty"`
	IsSystem    bool    `json:"is_system,omitempty"`
}

type RoleUpdateRequestDTO struct {
	RoleID      uuid.UUID `uri:"role_id" validate:"required,uuid"`
	Name        string    `json:"name" validate:"required"`
	Description *string   `json:"description,omitempty"`
	ScopeType   string    `json:"scope_type,omitempty"`
}

type RoleByIDRequestDTO struct {
	RoleID uuid.UUID `uri:"role_id" validate:"required,uuid"`
}

type RoleByKeyRequestDTO struct {
	Key string `uri:"key" validate:"required"`
}

func (r *RoleCreateRequestDTO) ToCreateCmd() roleAppCommands.CreateRoleCmd {
	cmd := roleAppCommands.CreateRoleCmd{
		Key:         r.Key,
		Name:        r.Name,
		Description: r.Description,
		IsSystem:    r.IsSystem,
	}

	// Parse scope type
	if r.ScopeType != "" {
		cmd.ScopeType = roleDomain.ScopeType(r.ScopeType)
	}

	return cmd
}

func (r *RoleUpdateRequestDTO) ToUpdateCmd() roleAppCommands.UpdateRoleCmd {
	cmd := roleAppCommands.UpdateRoleCmd{
		RoleID:      r.RoleID,
		Name:        r.Name,
		Description: r.Description,
	}

	// Parse scope type
	if r.ScopeType != "" {
		cmd.ScopeType = roleDomain.ScopeType(r.ScopeType)
	}

	return cmd
}
