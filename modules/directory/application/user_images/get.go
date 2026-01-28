package user_images

import (
	"context"
	userImageResult "nfxid/modules/directory/application/user_images/results"

	"github.com/google/uuid"
)

// GetUserImage 根据ID获取用户图片
func (s *Service) GetUserImage(ctx context.Context, userImageID uuid.UUID) (userImageResult.UserImageRO, error) {
	domainEntity, err := s.userImageRepo.Get.ByID(ctx, userImageID)
	if err != nil {
		return userImageResult.UserImageRO{}, err
	}
	return userImageResult.UserImageMapper(domainEntity), nil
}

// GetUserImagesByUserID 根据用户ID获取用户图片列表
func (s *Service) GetUserImagesByUserID(ctx context.Context, userID uuid.UUID) ([]userImageResult.UserImageRO, error) {
	domainEntities, err := s.userImageRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	results := make([]userImageResult.UserImageRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = userImageResult.UserImageMapper(entity)
	}
	return results, nil
}

// GetUserImagesByImageID 根据图片ID获取用户图片列表
func (s *Service) GetUserImagesByImageID(ctx context.Context, imageID uuid.UUID) ([]userImageResult.UserImageRO, error) {
	domainEntities, err := s.userImageRepo.Get.ByImageID(ctx, imageID)
	if err != nil {
		return nil, err
	}

	results := make([]userImageResult.UserImageRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = userImageResult.UserImageMapper(entity)
	}
	return results, nil
}

// GetCurrentUserImageByUserID 获取用户当前图片（display_order = 0）
func (s *Service) GetCurrentUserImageByUserID(ctx context.Context, userID uuid.UUID) (userImageResult.UserImageRO, error) {
	domainEntity, err := s.userImageRepo.Get.CurrentByUserID(ctx, userID)
	if err != nil {
		return userImageResult.UserImageRO{}, err
	}
	return userImageResult.UserImageMapper(domainEntity), nil
}
