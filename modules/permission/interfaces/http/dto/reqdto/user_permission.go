package reqdto

import (
	userPermissionCommands "nfxid/modules/permission/application/user_permission/commands"
	"github.com/google/uuid"
)

type UserPermissionAssignRequestDTO struct {
	UserID       uuid.UUID `json:"user_id" validate:"required,uuid"`
	PermissionID uuid.UUID `json:"permission_id" validate:"required,uuid"`
}

func (dto *UserPermissionAssignRequestDTO) ToAssignCmd() userPermissionCommands.AssignPermissionCmd {
	return userPermissionCommands.AssignPermissionCmd{
		UserID:       dto.UserID,
		PermissionID: dto.PermissionID,
	}
}

type UserPermissionRevokeRequestDTO struct {
	UserID       uuid.UUID `json:"user_id" validate:"required,uuid"`
	PermissionID uuid.UUID `json:"permission_id" validate:"required,uuid"`
}

func (dto *UserPermissionRevokeRequestDTO) ToRevokeCmd() userPermissionCommands.RevokePermissionCmd {
	return userPermissionCommands.RevokePermissionCmd{
		UserID:       dto.UserID,
		PermissionID: dto.PermissionID,
	}
}

type UserPermissionByUserIDRequestDTO struct {
	UserID uuid.UUID `params:"user_id" validate:"required,uuid"`
}

type UserPermissionCheckRequestDTO struct {
	UserID uuid.UUID `json:"user_id" validate:"required,uuid"`
	Tag    string    `json:"tag" validate:"required"`
}

func (dto *UserPermissionCheckRequestDTO) ToCheckCmd() userPermissionCommands.CheckPermissionCmd {
	return userPermissionCommands.CheckPermissionCmd{
		UserID: dto.UserID,
		Tag:    dto.Tag,
	}
}

