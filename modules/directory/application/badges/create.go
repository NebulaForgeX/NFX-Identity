package badges

import (
	"context"
	dirErr "nfxidentity/errors/src/directory"
	badgeCommands "nfxidentity/modules/directory/application/badges/commands"
	badgeDomain "nfxidentity/modules/directory/domain/badges"

	"github.com/google/uuid"
)

// CreateBadge 创建徽章
func (s *Service) CreateBadge(ctx context.Context, cmd badgeCommands.CreateBadgeCmd) (uuid.UUID, error) {
	// Check if badge name already exists
	if exists, _ := s.badgeRepo.Check.ByName(ctx, cmd.Name); exists {
		return uuid.Nil, dirErr.ErrNameAlreadyExists
	}

	// Create domain entity
	badge, err := badgeDomain.NewBadge(badgeDomain.NewBadgeParams{
		Name:        cmd.Name,
		Description: cmd.Description,
		IconURL:     cmd.IconURL,
		Color:       cmd.Color,
		Category:    cmd.Category,
		IsSystem:    cmd.IsSystem,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.badgeRepo.Create.New(ctx, badge); err != nil {
		return uuid.Nil, err
	}

	return badge.ID(), nil
}
