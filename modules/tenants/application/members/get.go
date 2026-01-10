package members

import (
	"context"
	memberResult "nfxid/modules/tenants/application/members/results"

	"github.com/google/uuid"
)

// GetMember 根据ID获取成员
func (s *Service) GetMember(ctx context.Context, memberID uuid.UUID) (memberResult.MemberRO, error) {
	domainEntity, err := s.memberRepo.Get.ByID(ctx, memberID)
	if err != nil {
		return memberResult.MemberRO{}, err
	}
	return memberResult.MemberMapper(domainEntity), nil
}

// GetMemberByMemberID 根据MemberID获取成员
func (s *Service) GetMemberByMemberID(ctx context.Context, memberID uuid.UUID) (memberResult.MemberRO, error) {
	domainEntity, err := s.memberRepo.Get.ByMemberID(ctx, memberID)
	if err != nil {
		return memberResult.MemberRO{}, err
	}
	return memberResult.MemberMapper(domainEntity), nil
}
