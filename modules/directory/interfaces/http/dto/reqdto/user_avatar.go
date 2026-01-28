package reqdto

import (
	userAvatarAppCommands "nfxid/modules/directory/application/user_avatars/commands"

	"github.com/google/uuid"
)

type UserAvatarCreateOrUpdateRequestDTO struct {
	UserID  uuid.UUID `json:"user_id" validate:"required,uuid"`
	ImageID uuid.UUID `json:"image_id" validate:"required,uuid"`
}

type UserAvatarUpdateImageIDRequestDTO struct {
	ImageID uuid.UUID `json:"image_id" validate:"required,uuid"`
}

func (r *UserAvatarCreateOrUpdateRequestDTO) ToCreateOrUpdateCmd() userAvatarAppCommands.CreateOrUpdateUserAvatarCmd {
	return userAvatarAppCommands.CreateOrUpdateUserAvatarCmd{
		UserID:  r.UserID,
		ImageID: r.ImageID,
	}
}

func (r *UserAvatarUpdateImageIDRequestDTO) ToUpdateImageIDCmd(userID uuid.UUID) userAvatarAppCommands.UpdateUserAvatarImageIDCmd {
	return userAvatarAppCommands.UpdateUserAvatarImageIDCmd{
		UserID:  userID,
		ImageID: r.ImageID,
	}
}
