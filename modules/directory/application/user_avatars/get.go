package user_avatars

import (
	"context"
	userAvatarResult "nfxid/modules/directory/application/user_avatars/results"

	"github.com/google/uuid"
)

// GetUserAvatarByUserID 根据用户ID获取用户头像
func (s *Service) GetUserAvatarByUserID(ctx context.Context, userID uuid.UUID) (userAvatarResult.UserAvatarRO, error) {
	domainEntity, err := s.userAvatarRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return userAvatarResult.UserAvatarRO{}, err
	}
	return userAvatarResult.UserAvatarMapper(domainEntity), nil
}

// GetUserAvatarByImageID 根据图片ID获取用户头像
func (s *Service) GetUserAvatarByImageID(ctx context.Context, imageID uuid.UUID) (userAvatarResult.UserAvatarRO, error) {
	domainEntity, err := s.userAvatarRepo.Get.ByImageID(ctx, imageID)
	if err != nil {
		return userAvatarResult.UserAvatarRO{}, err
	}
	return userAvatarResult.UserAvatarMapper(domainEntity), nil
}
