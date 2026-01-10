package member_app_roles

import (
	"context"
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
