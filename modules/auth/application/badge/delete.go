package badge

import (
	"context"

	"github.com/google/uuid"
)

type DeleteBadgeCmd struct {
	BadgeID uuid.UUID
}

func (s *Service) DeleteBadge(ctx context.Context, cmd DeleteBadgeCmd) error {
	b, err := s.badgeRepo.GetByID(ctx, cmd.BadgeID)
	if err != nil {
		return err
	}

	if err := b.Delete(); err != nil {
		return err
	}

	if err := s.badgeRepo.Update(ctx, b); err != nil {
		return err
	}

	return nil
}

