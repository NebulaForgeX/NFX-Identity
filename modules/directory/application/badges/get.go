package badges

import (
	"context"
	badgeResult "nfxid/modules/directory/application/badges/results"

	"github.com/google/uuid"
)

// GetBadge 根据ID获取徽章
func (s *Service) GetBadge(ctx context.Context, badgeID uuid.UUID) (badgeResult.BadgeRO, error) {
	domainEntity, err := s.badgeRepo.Get.ByID(ctx, badgeID)
	if err != nil {
		return badgeResult.BadgeRO{}, err
	}
	return badgeResult.BadgeMapper(domainEntity), nil
}

// GetBadgeByName 根据Name获取徽章
func (s *Service) GetBadgeByName(ctx context.Context, name string) (badgeResult.BadgeRO, error) {
	domainEntity, err := s.badgeRepo.Get.ByName(ctx, name)
	if err != nil {
		return badgeResult.BadgeRO{}, err
	}
	return badgeResult.BadgeMapper(domainEntity), nil
}

// GetAllBadges 获取所有徽章列表
func (s *Service) GetAllBadges(ctx context.Context, category *string, isSystem *bool) ([]badgeResult.BadgeRO, error) {
	// 注意：repository 层只有 ByCategory 方法，没有 GetAll 方法
	// 如果 category 为空，无法获取所有徽章，返回空列表
	if category == nil {
		return []badgeResult.BadgeRO{}, nil
	}
	
	domainEntities, err := s.badgeRepo.Get.ByCategory(ctx, *category)
	if err != nil {
		return nil, err
	}
	
	results := make([]badgeResult.BadgeRO, 0, len(domainEntities))
	for _, entity := range domainEntities {
		// 如果指定了isSystem，进行过滤
		if isSystem != nil {
			if entity.IsSystem() == *isSystem {
				results = append(results, badgeResult.BadgeMapper(entity))
			}
		} else {
			results = append(results, badgeResult.BadgeMapper(entity))
		}
	}
	return results, nil
}
