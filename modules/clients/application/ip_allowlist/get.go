package ip_allowlist

import (
	"context"
	ipAllowlistResult "nfxid/modules/clients/application/ip_allowlist/results"

	"github.com/google/uuid"
)

// GetIPAllowlist 根据ID获取IP白名单
func (s *Service) GetIPAllowlist(ctx context.Context, ipAllowlistID uuid.UUID) (ipAllowlistResult.IPAllowlistRO, error) {
	domainEntity, err := s.ipAllowlistRepo.Get.ByID(ctx, ipAllowlistID)
	if err != nil {
		return ipAllowlistResult.IPAllowlistRO{}, err
	}
	return ipAllowlistResult.IPAllowlistMapper(domainEntity), nil
}

// GetIPAllowlistByRuleID 根据RuleID获取IP白名单
func (s *Service) GetIPAllowlistByRuleID(ctx context.Context, ruleID string) (ipAllowlistResult.IPAllowlistRO, error) {
	domainEntity, err := s.ipAllowlistRepo.Get.ByRuleID(ctx, ruleID)
	if err != nil {
		return ipAllowlistResult.IPAllowlistRO{}, err
	}
	return ipAllowlistResult.IPAllowlistMapper(domainEntity), nil
}
