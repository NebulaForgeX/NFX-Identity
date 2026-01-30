package user_avatars

import (
	"context"
	"fmt"
	"strings"

	"nfxid/constants"
	"nfxid/events/directory"
	userAvatarCommands "nfxid/modules/directory/application/user_avatars/commands"
	userAvatarDomain "nfxid/modules/directory/domain/user_avatars"
	"nfxid/pkgs/kafkax/eventbus"
)

// CreateOrUpdateUserAvatar 创建或更新用户头像
// 流程：
// 1. 通过 gRPC 验证 Image 是否存在
// 2. 如果图片在 tmp 目录，移动到 avatar 目录
// 3. 若已有头像则先取旧 image_id，更新后再通过 gRPC 删除旧头像文件
// 4. 创建或更新用户头像关联
func (s *Service) CreateOrUpdateUserAvatar(ctx context.Context, cmd userAvatarCommands.CreateOrUpdateUserAvatarCmd) error {

	// 获取图片信息
	image, err := s.grpcClients.ImageClient.GetImageByID(ctx, cmd.ImageID.String())
	if err != nil {
		return fmt.Errorf("image not found: %w", err)
	}

	// 如果图片在 tmp 目录，移动到 avatar 目录
	if strings.Contains(image.StoragePath, constants.StoragePathTmp) {
		_, err := s.grpcClients.ImageClient.MoveImage(ctx, cmd.ImageID.String(), string(constants.ImageStorageTypeAvatar))
		if err != nil {
			return fmt.Errorf("failed to move image to avatar directory: %w", err)
		}
	}

	exists, _ := s.userAvatarRepo.Check.ByUserID(ctx, cmd.UserID)

	if exists {
		// 更新：先取旧头像的 image_id，更新后再删除旧图片文件
		oldAvatar, err := s.userAvatarRepo.Get.ByUserID(ctx, cmd.UserID)
		if err != nil {
			return err
		}
		oldImageID := oldAvatar.ImageID()

		if err := s.userAvatarRepo.Update.ImageID(ctx, cmd.UserID, cmd.ImageID); err != nil {
			return err
		}

		// 若旧头像与新头像不是同一张图，发布事件由 Image 服务删除旧图片
		if oldImageID != cmd.ImageID && s.busPublisher != nil {
			evt := directory.UserAvatarReplacedEvent{OldImageID: oldImageID.String()}
			_ = eventbus.PublishEvent(ctx, s.busPublisher, evt)
		}
		return nil
	}

	// Create new avatar
	userAvatar, err := userAvatarDomain.NewUserAvatar(userAvatarDomain.NewUserAvatarParams{
		UserID:  cmd.UserID,
		ImageID: cmd.ImageID,
	})
	if err != nil {
		return err
	}

	return s.userAvatarRepo.Create.New(ctx, userAvatar)
}
