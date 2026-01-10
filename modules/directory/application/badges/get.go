package badges

import (
	"context"
	badgeResult "nfxid/modules/directory/application/badges/results"

	"github.com/google/uuid"
)

// GetBadge 根据ID获取徽章
func (s *Service) GetBadge(ctx context.Context, badgeID uuid.UUID) (badgeResult.BadgeRO, error) {
	domainEntity, err := s.badgeRepo.Get.ByID(ctx, badgeID)
	if err != nil {
		return badgeResult.BadgeRO{}, err
	}
	return badgeResult.BadgeMapper(domainEntity), nil
}

// GetBadgeByName 根据Name获取徽章
func (s *Service) GetBadgeByName(ctx context.Context, name string) (badgeResult.BadgeRO, error) {
	domainEntity, err := s.badgeRepo.Get.ByName(ctx, name)
	if err != nil {
		return badgeResult.BadgeRO{}, err
	}
	return badgeResult.BadgeMapper(domainEntity), nil
}
