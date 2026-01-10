package member_roles

import (
	"context"
	memberRoleResult "nfxid/modules/tenants/application/member_roles/results"

	"github.com/google/uuid"
)

// GetMemberRole 根据ID获取成员角色
func (s *Service) GetMemberRole(ctx context.Context, memberRoleID uuid.UUID) (memberRoleResult.MemberRoleRO, error) {
	domainEntity, err := s.memberRoleRepo.Get.ByID(ctx, memberRoleID)
	if err != nil {
		return memberRoleResult.MemberRoleRO{}, err
	}
	return memberRoleResult.MemberRoleMapper(domainEntity), nil
}

// GetMemberRolesByMemberID 根据成员ID获取成员角色列表
func (s *Service) GetMemberRolesByMemberID(ctx context.Context, memberID uuid.UUID) ([]memberRoleResult.MemberRoleRO, error) {
	domainEntities, err := s.memberRoleRepo.Get.ByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	
	results := make([]memberRoleResult.MemberRoleRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = memberRoleResult.MemberRoleMapper(entity)
	}
	return results, nil
}
