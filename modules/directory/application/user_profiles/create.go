package user_profiles

import (
	"context"
	"time"
	userProfileCommands "nfxid/modules/directory/application/user_profiles/commands"
	userProfileDomain "nfxid/modules/directory/domain/user_profiles"

	"github.com/google/uuid"
)

// CreateUserProfile 创建用户资料
func (s *Service) CreateUserProfile(ctx context.Context, cmd userProfileCommands.CreateUserProfileCmd) (uuid.UUID, error) {
	// Check if user profile already exists
	if exists, _ := s.userProfileRepo.Check.ByUserID(ctx, cmd.UserID); exists {
		return uuid.Nil, userProfileDomain.ErrUserProfileAlreadyExists
	}

	var birthday *time.Time
	if cmd.Birthday != nil && *cmd.Birthday != "" {
		parsed, err := parseDateString(*cmd.Birthday)
		if err != nil {
			return uuid.Nil, err
		}
		birthday = parsed
	}

	// Create domain entity
	userProfile, err := userProfileDomain.NewUserProfile(userProfileDomain.NewUserProfileParams{
		UserID:       cmd.UserID,
		Role:         cmd.Role,
		FirstName:    cmd.FirstName,
		LastName:     cmd.LastName,
		Nickname:     cmd.Nickname,
		DisplayName:  cmd.DisplayName,
		AvatarID:     cmd.AvatarID,
		BackgroundID: cmd.BackgroundID,
		BackgroundIDs: cmd.BackgroundIDs,
		Bio:          cmd.Bio,
		Birthday:     birthday,
		Age:          cmd.Age,
		Gender:       cmd.Gender,
		Location:     cmd.Location,
		Website:      cmd.Website,
		Github:       cmd.Github,
		SocialLinks:  cmd.SocialLinks,
		Skills:       cmd.Skills,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.userProfileRepo.Create.New(ctx, userProfile); err != nil {
		return uuid.Nil, err
	}

	return userProfile.ID(), nil
}
