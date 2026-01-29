package user_images

import (
	"context"
	"fmt"
	"strings"

	"nfxid/constants"
	userImageCommands "nfxid/modules/directory/application/user_images/commands"
	userImageDomain "nfxid/modules/directory/domain/user_images"

	"github.com/google/uuid"
)

// CreateUserImage 创建用户图片
// 流程：
// 1. 通过 gRPC 验证 Image 是否存在
// 2. 如果图片在 tmp 目录，移动到 background 目录
// 3. 创建用户图片关联
func (s *Service) CreateUserImage(ctx context.Context, cmd userImageCommands.CreateUserImageCmd) (uuid.UUID, error) {
	// 通过 gRPC 验证 Image 是否存在
	if s.imageClient == nil {
		return uuid.Nil, fmt.Errorf("image client not configured")
	}

	// 获取图片信息
	image, err := s.imageClient.GetImageByID(ctx, cmd.ImageID.String())
	if err != nil {
		return uuid.Nil, fmt.Errorf("image not found: %w", err)
	}

	// 如果图片在 tmp 目录，移动到 background 目录
	if strings.Contains(image.StoragePath, constants.StoragePathTmp) {
		_, err := s.imageClient.MoveImage(ctx, cmd.ImageID.String(), string(constants.ImageStorageTypeBackground))
		if err != nil {
			return uuid.Nil, fmt.Errorf("failed to move image to background directory: %w", err)
		}
	}

	// Check if image already exists for this user
	if exists, _ := s.userImageRepo.Check.ByUserIDAndImageID(ctx, cmd.UserID, cmd.ImageID); exists {
		return uuid.Nil, userImageDomain.ErrUserImageNotFound // or return existing ID
	}

	// Create domain entity
	userImage, err := userImageDomain.NewUserImage(userImageDomain.NewUserImageParams{
		UserID:       cmd.UserID,
		ImageID:      cmd.ImageID,
		DisplayOrder: cmd.DisplayOrder,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.userImageRepo.Create.New(ctx, userImage); err != nil {
		return uuid.Nil, err
	}

	return userImage.ID(), nil
}
