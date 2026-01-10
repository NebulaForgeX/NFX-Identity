package user_occupations

import (
	"context"
	userOccupationDomain "nfxid/modules/directory/domain/user_occupations"
	userOccupationResult "nfxid/modules/directory/application/user_occupations/results"

	"github.com/google/uuid"
)

// GetUserOccupation 根据ID获取用户职业经历
func (s *Service) GetUserOccupation(ctx context.Context, userOccupationID uuid.UUID) (userOccupationResult.UserOccupationRO, error) {
	domainEntity, err := s.userOccupationRepo.Get.ByID(ctx, userOccupationID)
	if err != nil {
		return userOccupationResult.UserOccupationRO{}, err
	}
	return userOccupationResult.UserOccupationMapper(domainEntity), nil
}

// GetUserOccupationsByUserID 根据用户ID获取用户职业列表
func (s *Service) GetUserOccupationsByUserID(ctx context.Context, userID uuid.UUID, isCurrent *bool) ([]userOccupationResult.UserOccupationRO, error) {
	var domainEntities []*userOccupationDomain.UserOccupation
	var err error
	
	if isCurrent != nil && *isCurrent {
		domainEntities, err = s.userOccupationRepo.Get.CurrentByUserID(ctx, userID)
	} else {
		domainEntities, err = s.userOccupationRepo.Get.ByUserID(ctx, userID)
	}
	
	if err != nil {
		return nil, err
	}
	
	results := make([]userOccupationResult.UserOccupationRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = userOccupationResult.UserOccupationMapper(entity)
	}
	return results, nil
}
