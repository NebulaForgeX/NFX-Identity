package user_occupations

import (
	"context"
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
