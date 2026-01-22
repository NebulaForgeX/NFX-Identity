package images

import (
	"context"
	imageResult "nfxid/modules/image/application/images/results"

	"github.com/google/uuid"
)

// GetImage 根据ID获取图片
func (s *Service) GetImage(ctx context.Context, imageID uuid.UUID) (imageResult.ImageRO, error) {
	domainEntity, err := s.imageRepo.Get.ByID(ctx, imageID)
	if err != nil {
		return imageResult.ImageRO{}, err
	}
	return imageResult.ImageMapper(domainEntity), nil
}

// GetImagesByUserID 根据用户ID获取图片列表
func (s *Service) GetImagesByUserID(ctx context.Context, userID uuid.UUID) ([]imageResult.ImageRO, error) {
	domainEntities, err := s.imageRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	results := make([]imageResult.ImageRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = imageResult.ImageMapper(entity)
	}
	return results, nil
}

// GetImagesByTenantID 根据租户ID获取图片列表
func (s *Service) GetImagesByTenantID(ctx context.Context, tenantID uuid.UUID) ([]imageResult.ImageRO, error) {
	domainEntities, err := s.imageRepo.Get.ByTenantID(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	results := make([]imageResult.ImageRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = imageResult.ImageMapper(entity)
	}
	return results, nil
}

// GetImagesByTypeID 根据类型ID获取图片列表
func (s *Service) GetImagesByTypeID(ctx context.Context, typeID uuid.UUID) ([]imageResult.ImageRO, error) {
	domainEntities, err := s.imageRepo.Get.ByTypeID(ctx, typeID)
	if err != nil {
		return nil, err
	}

	results := make([]imageResult.ImageRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = imageResult.ImageMapper(entity)
	}
	return results, nil
}

// GetImagesBySourceDomain 根据源域名获取图片列表
func (s *Service) GetImagesBySourceDomain(ctx context.Context, sourceDomain string) ([]imageResult.ImageRO, error) {
	domainEntities, err := s.imageRepo.Get.BySourceDomain(ctx, sourceDomain)
	if err != nil {
		return nil, err
	}

	results := make([]imageResult.ImageRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = imageResult.ImageMapper(entity)
	}
	return results, nil
}

// GetPublicImagesByUserID 根据用户ID获取公开图片列表
func (s *Service) GetPublicImagesByUserID(ctx context.Context, userID uuid.UUID) ([]imageResult.ImageRO, error) {
	domainEntities, err := s.imageRepo.Get.PublicByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	results := make([]imageResult.ImageRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = imageResult.ImageMapper(entity)
	}
	return results, nil
}
