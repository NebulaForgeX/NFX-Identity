package reqdto

import (
	"nfxid/enums"
	permissionCommands "nfxid/modules/permission/application/permission/commands"

	"github.com/google/uuid"
)

type PermissionCreateRequestDTO struct {
	Tag         string `json:"tag" validate:"required,min=1,max=100"`
	Name        string `json:"name" validate:"required,min=1,max=255"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty" validate:"max=50"`
	IsSystem    bool   `json:"is_system,omitempty"`
}

func (dto *PermissionCreateRequestDTO) ToCreateCmd() permissionCommands.CreatePermissionCmd {
	return permissionCommands.CreatePermissionCmd{
		Tag:         dto.Tag,
		Name:        dto.Name,
		Description: dto.Description,
		Category:    enums.PermissionCategory(dto.Category), // Convert string to enum
		IsSystem:    dto.IsSystem,
	}
}

type PermissionUpdateRequestDTO struct {
	ID          uuid.UUID `params:"id" validate:"required,uuid"`
	Tag         string    `json:"tag" validate:"required,min=1,max=100"`
	Name        string    `json:"name" validate:"required,min=1,max=255"`
	Description string    `json:"description,omitempty"`
	Category    string    `json:"category,omitempty" validate:"max=50"`
}

func (dto *PermissionUpdateRequestDTO) ToUpdateCmd() permissionCommands.UpdatePermissionCmd {
	return permissionCommands.UpdatePermissionCmd{
		ID:          dto.ID,
		Tag:         dto.Tag,
		Name:        dto.Name,
		Description: dto.Description,
		Category:    enums.PermissionCategory(dto.Category), // Convert string to enum
	}
}

type PermissionByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type PermissionByTagRequestDTO struct {
	Tag string `params:"tag" validate:"required"`
}

type PermissionQueryParamsDTO struct {
	Category *string `query:"category"`
}
