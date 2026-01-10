package user_badges

import (
	"context"
	userBadgeResult "nfxid/modules/directory/application/user_badges/results"

	"github.com/google/uuid"
)

// GetUserBadge 根据ID获取用户徽章
func (s *Service) GetUserBadge(ctx context.Context, userBadgeID uuid.UUID) (userBadgeResult.UserBadgeRO, error) {
	domainEntity, err := s.userBadgeRepo.Get.ByID(ctx, userBadgeID)
	if err != nil {
		return userBadgeResult.UserBadgeRO{}, err
	}
	return userBadgeResult.UserBadgeMapper(domainEntity), nil
}
