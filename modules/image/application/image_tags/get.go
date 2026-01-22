package image_tags

import (
	"context"
	imageTagResult "nfxid/modules/image/application/image_tags/results"

	"github.com/google/uuid"
)

// GetImageTag 根据ID获取图片标签
func (s *Service) GetImageTag(ctx context.Context, imageTagID uuid.UUID) (imageTagResult.ImageTagRO, error) {
	domainEntity, err := s.imageTagRepo.Get.ByID(ctx, imageTagID)
	if err != nil {
		return imageTagResult.ImageTagRO{}, err
	}
	return imageTagResult.ImageTagMapper(domainEntity), nil
}

// GetImageTagsByImageID 根据图片ID获取图片标签列表
func (s *Service) GetImageTagsByImageID(ctx context.Context, imageID uuid.UUID) ([]imageTagResult.ImageTagRO, error) {
	domainEntities, err := s.imageTagRepo.Get.ByImageID(ctx, imageID)
	if err != nil {
		return nil, err
	}

	results := make([]imageTagResult.ImageTagRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = imageTagResult.ImageTagMapper(entity)
	}
	return results, nil
}

// GetImageTagsByTag 根据标签获取图片标签列表
func (s *Service) GetImageTagsByTag(ctx context.Context, tag string) ([]imageTagResult.ImageTagRO, error) {
	domainEntities, err := s.imageTagRepo.Get.ByTag(ctx, tag)
	if err != nil {
		return nil, err
	}

	results := make([]imageTagResult.ImageTagRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = imageTagResult.ImageTagMapper(entity)
	}
	return results, nil
}

// GetImageTagByImageIDAndTag 根据图片ID和标签获取图片标签
func (s *Service) GetImageTagByImageIDAndTag(ctx context.Context, imageID uuid.UUID, tag string) (imageTagResult.ImageTagRO, error) {
	domainEntity, err := s.imageTagRepo.Get.ByImageIDAndTag(ctx, imageID, tag)
	if err != nil {
		return imageTagResult.ImageTagRO{}, err
	}
	return imageTagResult.ImageTagMapper(domainEntity), nil
}
