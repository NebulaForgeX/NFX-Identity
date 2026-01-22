package image_variants

import (
	"context"
	imageVariantResult "nfxid/modules/image/application/image_variants/results"

	"github.com/google/uuid"
)

// GetImageVariant 根据ID获取图片变体
func (s *Service) GetImageVariant(ctx context.Context, imageVariantID uuid.UUID) (imageVariantResult.ImageVariantRO, error) {
	domainEntity, err := s.imageVariantRepo.Get.ByID(ctx, imageVariantID)
	if err != nil {
		return imageVariantResult.ImageVariantRO{}, err
	}
	return imageVariantResult.ImageVariantMapper(domainEntity), nil
}

// GetImageVariantsByImageID 根据图片ID获取图片变体列表
func (s *Service) GetImageVariantsByImageID(ctx context.Context, imageID uuid.UUID) ([]imageVariantResult.ImageVariantRO, error) {
	domainEntities, err := s.imageVariantRepo.Get.ByImageID(ctx, imageID)
	if err != nil {
		return nil, err
	}

	results := make([]imageVariantResult.ImageVariantRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = imageVariantResult.ImageVariantMapper(entity)
	}
	return results, nil
}

// GetImageVariantByImageIDAndVariantKey 根据图片ID和变体Key获取图片变体
func (s *Service) GetImageVariantByImageIDAndVariantKey(ctx context.Context, imageID uuid.UUID, variantKey string) (imageVariantResult.ImageVariantRO, error) {
	domainEntity, err := s.imageVariantRepo.Get.ByImageIDAndVariantKey(ctx, imageID, variantKey)
	if err != nil {
		return imageVariantResult.ImageVariantRO{}, err
	}
	return imageVariantResult.ImageVariantMapper(domainEntity), nil
}
