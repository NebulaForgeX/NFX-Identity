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
