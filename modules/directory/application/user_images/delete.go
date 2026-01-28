package user_images

import (
	"context"
	userImageCommands "nfxid/modules/directory/application/user_images/commands"
)

// DeleteUserImage 删除用户图片
func (s *Service) DeleteUserImage(ctx context.Context, cmd userImageCommands.DeleteUserImageCmd) error {
	// Get domain entity
	userImage, err := s.userImageRepo.Get.ByID(ctx, cmd.UserImageID)
	if err != nil {
		return err
	}

	// Delete domain entity
	if err := userImage.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.userImageRepo.Update.Generic(ctx, userImage)
}
