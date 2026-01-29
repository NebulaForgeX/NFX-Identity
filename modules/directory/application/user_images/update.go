package user_images

import (
	"context"
	userImageCommands "nfxid/modules/directory/application/user_images/commands"
)

// UpdateUserImageDisplayOrder 更新用户图片显示顺序
func (s *Service) UpdateUserImageDisplayOrder(ctx context.Context, cmd userImageCommands.UpdateUserImageDisplayOrderCmd) error {
	// Get domain entity
	userImage, err := s.userImageRepo.Get.ByID(ctx, cmd.UserImageID)
	if err != nil {
		return err
	}

	// Update domain entity
	if err := userImage.UpdateDisplayOrder(cmd.DisplayOrder); err != nil {
		return err
	}

	// Save to repository
	return s.userImageRepo.Update.Generic(ctx, userImage)
}

// UpdateUserImageImageID 更新用户图片ID
func (s *Service) UpdateUserImageImageID(ctx context.Context, cmd userImageCommands.UpdateUserImageImageIDCmd) error {
	// update 时通过 gRPC 问 Image 服务该 image 是否存在
	if s.imageClient != nil {
		if _, err := s.imageClient.GetImageByID(ctx, cmd.ImageID.String()); err != nil {
			return err
		}
	}

	// Get domain entity
	userImage, err := s.userImageRepo.Get.ByID(ctx, cmd.UserImageID)
	if err != nil {
		return err
	}

	// Update domain entity
	if err := userImage.UpdateImageID(cmd.ImageID); err != nil {
		return err
	}

	// Save to repository
	return s.userImageRepo.Update.Generic(ctx, userImage)
}
