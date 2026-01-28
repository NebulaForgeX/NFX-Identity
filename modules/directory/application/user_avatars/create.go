package user_avatars

import (
	"context"
	userAvatarCommands "nfxid/modules/directory/application/user_avatars/commands"
	userAvatarDomain "nfxid/modules/directory/domain/user_avatars"
)

// CreateOrUpdateUserAvatar 创建或更新用户头像
func (s *Service) CreateOrUpdateUserAvatar(ctx context.Context, cmd userAvatarCommands.CreateOrUpdateUserAvatarCmd) error {
	// Check if avatar already exists
	exists, _ := s.userAvatarRepo.Check.ByUserID(ctx, cmd.UserID)
	
	if exists {
		// Update existing avatar
		return s.userAvatarRepo.Update.ImageID(ctx, cmd.UserID, cmd.ImageID)
	}

	// Create new avatar
	userAvatar, err := userAvatarDomain.NewUserAvatar(userAvatarDomain.NewUserAvatarParams{
		UserID:  cmd.UserID,
		ImageID: cmd.ImageID,
	})
	if err != nil {
		return err
	}

	// Save to repository
	return s.userAvatarRepo.Create.New(ctx, userAvatar)
}
