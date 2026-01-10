package member_app_roles

import (
	"context"
	memberAppRoleDomain "nfxid/modules/tenants/domain/member_app_roles"
	memberAppRoleResult "nfxid/modules/tenants/application/member_app_roles/results"

	"github.com/google/uuid"
)

// GetMemberAppRole 根据ID获取成员应用角色
func (s *Service) GetMemberAppRole(ctx context.Context, memberAppRoleID uuid.UUID) (memberAppRoleResult.MemberAppRoleRO, error) {
	domainEntity, err := s.memberAppRoleRepo.Get.ByID(ctx, memberAppRoleID)
	if err != nil {
		return memberAppRoleResult.MemberAppRoleRO{}, err
	}
	return memberAppRoleResult.MemberAppRoleMapper(domainEntity), nil
}

// GetMemberAppRolesByMemberID 根据成员ID获取成员应用角色列表
func (s *Service) GetMemberAppRolesByMemberID(ctx context.Context, memberID uuid.UUID, appID *uuid.UUID) ([]memberAppRoleResult.MemberAppRoleRO, error) {
	var domainEntities []*memberAppRoleDomain.MemberAppRole
	var err error
	
	if appID != nil {
		domainEntities, err = s.memberAppRoleRepo.Get.ByMemberIDAndAppID(ctx, memberID, *appID)
	} else {
		domainEntities, err = s.memberAppRoleRepo.Get.ByMemberID(ctx, memberID)
	}
	
	if err != nil {
		return nil, err
	}
	
	results := make([]memberAppRoleResult.MemberAppRoleRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = memberAppRoleResult.MemberAppRoleMapper(entity)
	}
	return results, nil
}
