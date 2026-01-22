package image_types

import (
	"context"
	imageTypeResult "nfxid/modules/image/application/image_types/results"

	"github.com/google/uuid"
)

// GetImageType 根据ID获取图片类型
func (s *Service) GetImageType(ctx context.Context, imageTypeID uuid.UUID) (imageTypeResult.ImageTypeRO, error) {
	domainEntity, err := s.imageTypeRepo.Get.ByID(ctx, imageTypeID)
	if err != nil {
		return imageTypeResult.ImageTypeRO{}, err
	}
	return imageTypeResult.ImageTypeMapper(domainEntity), nil
}

// GetImageTypeByKey 根据Key获取图片类型
func (s *Service) GetImageTypeByKey(ctx context.Context, key string) (imageTypeResult.ImageTypeRO, error) {
	domainEntity, err := s.imageTypeRepo.Get.ByKey(ctx, key)
	if err != nil {
		return imageTypeResult.ImageTypeRO{}, err
	}
	return imageTypeResult.ImageTypeMapper(domainEntity), nil
}

// GetAllImageTypes 获取所有图片类型
func (s *Service) GetAllImageTypes(ctx context.Context) ([]imageTypeResult.ImageTypeRO, error) {
	domainEntities, err := s.imageTypeRepo.Get.All(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]imageTypeResult.ImageTypeRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = imageTypeResult.ImageTypeMapper(entity)
	}
	return results, nil
}

// GetSystemImageTypes 获取系统图片类型
func (s *Service) GetSystemImageTypes(ctx context.Context) ([]imageTypeResult.ImageTypeRO, error) {
	domainEntities, err := s.imageTypeRepo.Get.SystemTypes(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]imageTypeResult.ImageTypeRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = imageTypeResult.ImageTypeMapper(entity)
	}
	return results, nil
}
