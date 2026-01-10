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
