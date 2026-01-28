package user_images

import (
	"context"
	userImageCommands "nfxid/modules/directory/application/user_images/commands"
	userImageDomain "nfxid/modules/directory/domain/user_images"

	"github.com/google/uuid"
)

// CreateUserImage 创建用户图片
func (s *Service) CreateUserImage(ctx context.Context, cmd userImageCommands.CreateUserImageCmd) (uuid.UUID, error) {
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
