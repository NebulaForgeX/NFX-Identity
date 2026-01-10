package domain_verifications

import (
	"context"
	domainVerificationDomain "nfxid/modules/tenants/domain/domain_verifications"
	domainVerificationResult "nfxid/modules/tenants/application/domain_verifications/results"

	"github.com/google/uuid"
)

// GetDomainVerification 根据ID获取域名验证
func (s *Service) GetDomainVerification(ctx context.Context, domainVerificationID uuid.UUID) (domainVerificationResult.DomainVerificationRO, error) {
	domainEntity, err := s.domainVerificationRepo.Get.ByID(ctx, domainVerificationID)
	if err != nil {
		return domainVerificationResult.DomainVerificationRO{}, err
	}
	return domainVerificationResult.DomainVerificationMapper(domainEntity), nil
}

// GetDomainVerificationByDomain 根据域名获取域名验证
func (s *Service) GetDomainVerificationByDomain(ctx context.Context, domain string) (domainVerificationResult.DomainVerificationRO, error) {
	domainEntity, err := s.domainVerificationRepo.Get.ByDomain(ctx, domain)
	if err != nil {
		return domainVerificationResult.DomainVerificationRO{}, err
	}
	return domainVerificationResult.DomainVerificationMapper(domainEntity), nil
}

// GetDomainVerificationsByTenantID 根据租户ID获取域名验证列表
func (s *Service) GetDomainVerificationsByTenantID(ctx context.Context, tenantID uuid.UUID, status *domainVerificationDomain.VerificationStatus) ([]domainVerificationResult.DomainVerificationRO, error) {
	domainEntities, err := s.domainVerificationRepo.Get.ByTenantID(ctx, tenantID)
	if err != nil {
		return nil, err
	}
	
	results := make([]domainVerificationResult.DomainVerificationRO, 0, len(domainEntities))
	for _, entity := range domainEntities {
		// 如果指定了status，进行过滤
		if status != nil {
			if entity.Status() == *status {
				results = append(results, domainVerificationResult.DomainVerificationMapper(entity))
			}
		} else {
			results = append(results, domainVerificationResult.DomainVerificationMapper(entity))
		}
	}
	return results, nil
}
