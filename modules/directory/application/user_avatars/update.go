package user_avatars

import (
	"context"
	userAvatarCommands "nfxid/modules/directory/application/user_avatars/commands"
)

// UpdateUserAvatarImageID 更新用户头像图片ID
func (s *Service) UpdateUserAvatarImageID(ctx context.Context, cmd userAvatarCommands.UpdateUserAvatarImageIDCmd) error {
	// update 时通过 gRPC 问 Image 服务该 image 是否存在（wiring 未配 ImageAddr 时 ImageClient 为 nil，跳过校验）
	if s.grpcClients.ImageClient != nil {
		if _, err := s.grpcClients.ImageClient.GetImageByID(ctx, cmd.ImageID.String()); err != nil {
			return err
		}
	}

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
