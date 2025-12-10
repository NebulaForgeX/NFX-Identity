package profile_badge

import (
	"context"

	"github.com/google/uuid"
)

type DeleteProfileBadgeCmd struct {
	ProfileBadgeID uuid.UUID
}

func (s *Service) DeleteProfileBadge(ctx context.Context, cmd DeleteProfileBadgeCmd) error {
	return s.profileBadgeRepo.Delete(ctx, cmd.ProfileBadgeID)
}

type DeleteProfileBadgeByProfileAndBadgeCmd struct {
	ProfileID uuid.UUID
	BadgeID   uuid.UUID
}

func (s *Service) DeleteProfileBadgeByProfileAndBadge(ctx context.Context, cmd DeleteProfileBadgeByProfileAndBadgeCmd) error {
	return s.profileBadgeRepo.DeleteByProfileAndBadge(ctx, cmd.ProfileID, cmd.BadgeID)
}

