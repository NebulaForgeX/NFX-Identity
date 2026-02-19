package reqdto

import (
	permissionAppCommands "nfxid/modules/access/application/permissions/commands"

	"github.com/google/uuid"
)

type PermissionCreateRequestDTO struct {
	Key         string  `json:"key" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	IsSystem    bool    `json:"is_system,omitempty"`
}

type PermissionUpdateRequestDTO struct {
	ID          uuid.UUID `uri:"id" validate:"required,uuid"`
	Name        string    `json:"name" validate:"required"`
	Description *string   `json:"description,omitempty"`
}

type PermissionByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type PermissionByKeyRequestDTO struct {
	Key string `uri:"key" validate:"required"`
}

func (r *PermissionCreateRequestDTO) ToCreateCmd() permissionAppCommands.CreatePermissionCmd {
	return permissionAppCommands.CreatePermissionCmd{
		Key:         r.Key,
		Name:        r.Name,
		Description: r.Description,
		IsSystem:    r.IsSystem,
	}
}

func (r *PermissionUpdateRequestDTO) ToUpdateCmd() permissionAppCommands.UpdatePermissionCmd {
	return permissionAppCommands.UpdatePermissionCmd{
		PermissionID: r.ID,
		Name:         r.Name,
		Description:  r.Description,
	}
}
