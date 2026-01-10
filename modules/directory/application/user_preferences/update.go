package user_preferences

import (
	"context"
	userPreferenceCommands "nfxid/modules/directory/application/user_preferences/commands"
)

// UpdateUserPreference 更新用户偏好
func (s *Service) UpdateUserPreference(ctx context.Context, cmd userPreferenceCommands.UpdateUserPreferenceCmd) error {
	// Get domain entity
	userPreference, err := s.userPreferenceRepo.Get.ByID(ctx, cmd.UserPreferenceID)
	if err != nil {
		return err
	}

	// Update domain entity
	if err := userPreference.Update(cmd.Theme, cmd.Language, cmd.Timezone, cmd.Notifications, cmd.Privacy, cmd.Display, cmd.Other); err != nil {
		return err
	}

	// Save to repository
	return s.userPreferenceRepo.Update.Generic(ctx, userPreference)
}
