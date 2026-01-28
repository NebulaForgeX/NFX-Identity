package user_avatars

import (
	"context"
	userAvatarCommands "nfxid/modules/directory/application/user_avatars/commands"
)

// UpdateUserAvatarImageID 更新用户头像图片ID
func (s *Service) UpdateUserAvatarImageID(ctx context.Context, cmd userAvatarCommands.UpdateUserAvatarImageIDCmd) error {
	// Get domain entity
	userAvatar, err := s.userAvatarRepo.Get.ByUserID(ctx, cmd.UserID)
	if err != nil {
		return err
	}

	// Update domain entity
	if err := userAvatar.UpdateImageID(cmd.ImageID); err != nil {
		return err
	}

	// Save to repository
	return s.userAvatarRepo.Update.Generic(ctx, userAvatar)
}
