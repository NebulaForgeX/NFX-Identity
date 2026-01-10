package user_profiles

import (
	"context"
	"time"
	userProfileCommands "nfxid/modules/directory/application/user_profiles/commands"
)

// UpdateUserProfile 更新用户资料
func (s *Service) UpdateUserProfile(ctx context.Context, cmd userProfileCommands.UpdateUserProfileCmd) error {
	// Get domain entity
	userProfile, err := s.userProfileRepo.Get.ByID(ctx, cmd.UserProfileID)
	if err != nil {
		return err
	}

	var birthday *time.Time
	if cmd.Birthday != nil && *cmd.Birthday != "" {
		parsed, err := time.Parse(time.RFC3339, *cmd.Birthday)
		if err != nil {
			return err
		}
		birthday = &parsed
	}

	// Update domain entity
	if err := userProfile.Update(cmd.Role, cmd.FirstName, cmd.LastName, cmd.Nickname, cmd.DisplayName, cmd.AvatarID, cmd.BackgroundID, cmd.BackgroundIDs, cmd.Bio, cmd.Gender, cmd.Location, cmd.Website, cmd.Github, birthday, cmd.Age, cmd.SocialLinks, cmd.Skills); err != nil {
		return err
	}

	// Save to repository
	return s.userProfileRepo.Update.Generic(ctx, userProfile)
}
