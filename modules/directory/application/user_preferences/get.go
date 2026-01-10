package user_preferences

import (
	"context"
	userPreferenceResult "nfxid/modules/directory/application/user_preferences/results"

	"github.com/google/uuid"
)

// GetUserPreference 根据ID获取用户偏好
func (s *Service) GetUserPreference(ctx context.Context, userPreferenceID uuid.UUID) (userPreferenceResult.UserPreferenceRO, error) {
	domainEntity, err := s.userPreferenceRepo.Get.ByID(ctx, userPreferenceID)
	if err != nil {
		return userPreferenceResult.UserPreferenceRO{}, err
	}
	return userPreferenceResult.UserPreferenceMapper(domainEntity), nil
}

// GetUserPreferenceByUserID 根据UserID获取用户偏好
func (s *Service) GetUserPreferenceByUserID(ctx context.Context, userID uuid.UUID) (userPreferenceResult.UserPreferenceRO, error) {
	domainEntity, err := s.userPreferenceRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return userPreferenceResult.UserPreferenceRO{}, err
	}
	return userPreferenceResult.UserPreferenceMapper(domainEntity), nil
}
