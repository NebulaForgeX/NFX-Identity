package domain_verifications

import (
	"context"
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
