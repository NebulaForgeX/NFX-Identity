package user_profiles

import (
	"context"
	userProfileResult "nfxid/modules/directory/application/user_profiles/results"

	"github.com/google/uuid"
)

// GetUserProfile 根据ID获取用户资料
func (s *Service) GetUserProfile(ctx context.Context, userProfileID uuid.UUID) (userProfileResult.UserProfileRO, error) {
	domainEntity, err := s.userProfileRepo.Get.ByID(ctx, userProfileID)
	if err != nil {
		return userProfileResult.UserProfileRO{}, err
	}
	return userProfileResult.UserProfileMapper(domainEntity), nil
}

// GetUserProfileByUserID 根据UserID获取用户资料
func (s *Service) GetUserProfileByUserID(ctx context.Context, userID uuid.UUID) (userProfileResult.UserProfileRO, error) {
	domainEntity, err := s.userProfileRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return userProfileResult.UserProfileRO{}, err
	}
	return userProfileResult.UserProfileMapper(domainEntity), nil
}
