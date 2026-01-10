package user_profiles

import (
	"context"
	userProfileCommands "nfxid/modules/directory/application/user_profiles/commands"
)

// DeleteUserProfile 删除用户资料（软删除）
func (s *Service) DeleteUserProfile(ctx context.Context, cmd userProfileCommands.DeleteUserProfileCmd) error {
	// Get domain entity
	userProfile, err := s.userProfileRepo.Get.ByID(ctx, cmd.UserProfileID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := userProfile.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.userProfileRepo.Update.Generic(ctx, userProfile)
}
