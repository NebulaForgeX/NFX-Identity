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

// GetUserBadgesByUserID 根据用户ID获取用户徽章列表
func (s *Service) GetUserBadgesByUserID(ctx context.Context, userID uuid.UUID) ([]userBadgeResult.UserBadgeRO, error) {
	domainEntities, err := s.userBadgeRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	
	results := make([]userBadgeResult.UserBadgeRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = userBadgeResult.UserBadgeMapper(entity)
	}
	return results, nil
}

// GetUserBadgesByBadgeID 根据徽章ID获取用户徽章列表
func (s *Service) GetUserBadgesByBadgeID(ctx context.Context, badgeID uuid.UUID) ([]userBadgeResult.UserBadgeRO, error) {
	domainEntities, err := s.userBadgeRepo.Get.ByBadgeID(ctx, badgeID)
	if err != nil {
		return nil, err
	}
	
	results := make([]userBadgeResult.UserBadgeRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = userBadgeResult.UserBadgeMapper(entity)
	}
	return results, nil
}
