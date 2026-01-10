package member_groups

import (
	"context"
	memberGroupResult "nfxid/modules/tenants/application/member_groups/results"

	"github.com/google/uuid"
)

// GetMemberGroup 根据ID获取成员组
func (s *Service) GetMemberGroup(ctx context.Context, memberGroupID uuid.UUID) (memberGroupResult.MemberGroupRO, error) {
	domainEntity, err := s.memberGroupRepo.Get.ByID(ctx, memberGroupID)
	if err != nil {
		return memberGroupResult.MemberGroupRO{}, err
	}
	return memberGroupResult.MemberGroupMapper(domainEntity), nil
}

// GetMemberGroupsByMemberID 根据成员ID获取成员组列表
func (s *Service) GetMemberGroupsByMemberID(ctx context.Context, memberID uuid.UUID) ([]memberGroupResult.MemberGroupRO, error) {
	domainEntities, err := s.memberGroupRepo.Get.ByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	
	results := make([]memberGroupResult.MemberGroupRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = memberGroupResult.MemberGroupMapper(entity)
	}
	return results, nil
}

// GetMemberGroupsByGroupID 根据组ID获取成员组列表
func (s *Service) GetMemberGroupsByGroupID(ctx context.Context, groupID uuid.UUID) ([]memberGroupResult.MemberGroupRO, error) {
	domainEntities, err := s.memberGroupRepo.Get.ByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	
	results := make([]memberGroupResult.MemberGroupRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = memberGroupResult.MemberGroupMapper(entity)
	}
	return results, nil
}
