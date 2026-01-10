package user_educations

import (
	"context"
	userEducationResult "nfxid/modules/directory/application/user_educations/results"

	"github.com/google/uuid"
)

// GetUserEducation 根据ID获取用户教育经历
func (s *Service) GetUserEducation(ctx context.Context, userEducationID uuid.UUID) (userEducationResult.UserEducationRO, error) {
	domainEntity, err := s.userEducationRepo.Get.ByID(ctx, userEducationID)
	if err != nil {
		return userEducationResult.UserEducationRO{}, err
	}
	return userEducationResult.UserEducationMapper(domainEntity), nil
}

// GetUserEducationsByUserID 根据用户ID获取用户教育列表
func (s *Service) GetUserEducationsByUserID(ctx context.Context, userID uuid.UUID) ([]userEducationResult.UserEducationRO, error) {
	domainEntities, err := s.userEducationRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	
	results := make([]userEducationResult.UserEducationRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = userEducationResult.UserEducationMapper(entity)
	}
	return results, nil
}
