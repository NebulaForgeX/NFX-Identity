package user_preferences

import (
	"context"
	userPreferenceCommands "nfxid/modules/directory/application/user_preferences/commands"
)

// DeleteUserPreference 删除用户偏好（软删除）
func (s *Service) DeleteUserPreference(ctx context.Context, cmd userPreferenceCommands.DeleteUserPreferenceCmd) error {
	// Get domain entity
	userPreference, err := s.userPreferenceRepo.Get.ByID(ctx, cmd.UserPreferenceID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := userPreference.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.userPreferenceRepo.Update.Generic(ctx, userPreference)
}
