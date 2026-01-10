package members

import (
	"context"
	memberDomain "nfxid/modules/tenants/domain/members"
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

// GetMemberByUserID 根据用户ID和租户ID获取成员
func (s *Service) GetMemberByUserID(ctx context.Context, userID, tenantID uuid.UUID) (memberResult.MemberRO, error) {
	domainEntity, err := s.memberRepo.Get.ByTenantIDAndUserID(ctx, tenantID, userID)
	if err != nil {
		return memberResult.MemberRO{}, err
	}
	return memberResult.MemberMapper(domainEntity), nil
}

// GetMembersByTenantID 根据租户ID获取成员列表
func (s *Service) GetMembersByTenantID(ctx context.Context, tenantID uuid.UUID, status *memberDomain.MemberStatus) ([]memberResult.MemberRO, error) {
	var domainEntities []*memberDomain.Member
	var err error
	
	if status != nil {
		domainEntities, err = s.memberRepo.Get.ByTenantIDAndStatus(ctx, tenantID, *status)
	} else {
		domainEntities, err = s.memberRepo.Get.ByTenantID(ctx, tenantID)
	}
	
	if err != nil {
		return nil, err
	}
	
	results := make([]memberResult.MemberRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = memberResult.MemberMapper(entity)
	}
	return results, nil
}
