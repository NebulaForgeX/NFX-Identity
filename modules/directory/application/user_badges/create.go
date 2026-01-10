package user_badges

import (
	"context"
	"time"
	userBadgeCommands "nfxid/modules/directory/application/user_badges/commands"
	userBadgeDomain "nfxid/modules/directory/domain/user_badges"

	"github.com/google/uuid"
)

// CreateUserBadge 创建用户徽章
func (s *Service) CreateUserBadge(ctx context.Context, cmd userBadgeCommands.CreateUserBadgeCmd) (uuid.UUID, error) {
	// Check if user badge already exists
	if exists, _ := s.userBadgeRepo.Check.ByUserIDAndBadgeID(ctx, cmd.UserID, cmd.BadgeID); exists {
		return uuid.Nil, userBadgeDomain.ErrUserBadgeAlreadyExists
	}

	var earnedAt *time.Time
	if cmd.EarnedAt != nil && *cmd.EarnedAt != "" {
		parsed, err := time.Parse(time.RFC3339, *cmd.EarnedAt)
		if err != nil {
			return uuid.Nil, err
		}
		earnedAt = &parsed
	}

	// Create domain entity
	userBadge, err := userBadgeDomain.NewUserBadge(userBadgeDomain.NewUserBadgeParams{
		UserID:      cmd.UserID,
		BadgeID:     cmd.BadgeID,
		Description: cmd.Description,
		Level:       cmd.Level,
		EarnedAt:    earnedAt,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.userBadgeRepo.Create.New(ctx, userBadge); err != nil {
		return uuid.Nil, err
	}

	return userBadge.ID(), nil
}
