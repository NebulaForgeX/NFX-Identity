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
