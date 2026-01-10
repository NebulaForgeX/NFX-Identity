package user_preferences

import (
	"context"
	userPreferenceCommands "nfxid/modules/directory/application/user_preferences/commands"
	userPreferenceDomain "nfxid/modules/directory/domain/user_preferences"

	"github.com/google/uuid"
)

// CreateUserPreference 创建用户偏好
func (s *Service) CreateUserPreference(ctx context.Context, cmd userPreferenceCommands.CreateUserPreferenceCmd) (uuid.UUID, error) {
	// Check if user preference already exists
	if exists, _ := s.userPreferenceRepo.Check.ByUserID(ctx, cmd.UserID); exists {
		return uuid.Nil, userPreferenceDomain.ErrUserPreferenceAlreadyExists
	}

	// Create domain entity
	userPreference, err := userPreferenceDomain.NewUserPreference(userPreferenceDomain.NewUserPreferenceParams{
		UserID:        cmd.UserID,
		Theme:        cmd.Theme,
		Language:     cmd.Language,
		Timezone:     cmd.Timezone,
		Notifications: cmd.Notifications,
		Privacy:      cmd.Privacy,
		Display:      cmd.Display,
		Other:        cmd.Other,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.userPreferenceRepo.Create.New(ctx, userPreference); err != nil {
		return uuid.Nil, err
	}

	return userPreference.ID(), nil
}
