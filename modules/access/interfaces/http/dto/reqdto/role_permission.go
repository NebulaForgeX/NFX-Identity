package reqdto

import (
	rolePermissionAppCommands "nfxid/modules/access/application/role_permissions/commands"

	"github.com/google/uuid"
)

type RolePermissionCreateRequestDTO struct {
	RoleID       uuid.UUID  `json:"role_id" validate:"required"`
	PermissionID uuid.UUID  `json:"permission_id" validate:"required"`
	CreatedBy    *uuid.UUID `json:"created_by,omitempty"`
}

type RolePermissionByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type RolePermissionByRoleAndPermissionRequestDTO struct {
	RoleID       uuid.UUID `json:"role_id" validate:"required"`
	PermissionID uuid.UUID `json:"permission_id" validate:"required"`
}

type RolePermissionByRoleIDRequestDTO struct {
	RoleID uuid.UUID `params:"role_id" validate:"required,uuid"`
}

func (r *RolePermissionCreateRequestDTO) ToCreateCmd() rolePermissionAppCommands.CreateRolePermissionCmd {
	return rolePermissionAppCommands.CreateRolePermissionCmd{
		RoleID:       r.RoleID,
		PermissionID: r.PermissionID,
		CreatedBy:    r.CreatedBy,
	}
}
