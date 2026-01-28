package reqdto

import (
	userImageAppCommands "nfxid/modules/directory/application/user_images/commands"

	"github.com/google/uuid"
)

type UserImageCreateRequestDTO struct {
	UserID       uuid.UUID `json:"user_id" validate:"required,uuid"`
	ImageID      uuid.UUID `json:"image_id" validate:"required,uuid"`
	DisplayOrder int       `json:"display_order,omitempty"`
}

type UserImageUpdateDisplayOrderRequestDTO struct {
	DisplayOrder int `json:"display_order" validate:"required"`
}

type UserImageUpdateImageIDRequestDTO struct {
	ImageID uuid.UUID `json:"image_id" validate:"required,uuid"`
}

func (r *UserImageCreateRequestDTO) ToCreateCmd() userImageAppCommands.CreateUserImageCmd {
	return userImageAppCommands.CreateUserImageCmd{
		UserID:       r.UserID,
		ImageID:      r.ImageID,
		DisplayOrder: r.DisplayOrder,
	}
}

func (r *UserImageUpdateDisplayOrderRequestDTO) ToUpdateDisplayOrderCmd(userImageID uuid.UUID) userImageAppCommands.UpdateUserImageDisplayOrderCmd {
	return userImageAppCommands.UpdateUserImageDisplayOrderCmd{
		UserImageID: userImageID,
		DisplayOrder: r.DisplayOrder,
	}
}

func (r *UserImageUpdateImageIDRequestDTO) ToUpdateImageIDCmd(userImageID uuid.UUID) userImageAppCommands.UpdateUserImageImageIDCmd {
	return userImageAppCommands.UpdateUserImageImageIDCmd{
		UserImageID: userImageID,
		ImageID:     r.ImageID,
	}
}
