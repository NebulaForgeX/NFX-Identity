package groups

import (
	"context"
	groupResult "nfxid/modules/tenants/application/groups/results"

	"github.com/google/uuid"
)

// GetGroup 根据ID获取组
func (s *Service) GetGroup(ctx context.Context, groupID uuid.UUID) (groupResult.GroupRO, error) {
	domainEntity, err := s.groupRepo.Get.ByID(ctx, groupID)
	if err != nil {
		return groupResult.GroupRO{}, err
	}
	return groupResult.GroupMapper(domainEntity), nil
}

// GetGroupsByTenantID 根据租户ID获取组列表
func (s *Service) GetGroupsByTenantID(ctx context.Context, tenantID uuid.UUID, parentID *uuid.UUID) ([]groupResult.GroupRO, error) {
	domainEntities, err := s.groupRepo.Get.ByTenantID(ctx, tenantID)
	if err != nil {
		return nil, err
	}
	
	results := make([]groupResult.GroupRO, 0, len(domainEntities))
	for _, entity := range domainEntities {
		// 如果指定了parentID，进行过滤
		if parentID != nil {
			if entity.ParentGroupID() != nil && *entity.ParentGroupID() == *parentID {
				results = append(results, groupResult.GroupMapper(entity))
			}
		} else {
			results = append(results, groupResult.GroupMapper(entity))
		}
	}
	return results, nil
}
