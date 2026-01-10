package domain_verifications

import (
	domainVerificationDomain "nfxid/modules/tenants/domain/domain_verifications"
)

type Service struct {
	domainVerificationRepo *domainVerificationDomain.Repo
}

func NewService(
	domainVerificationRepo *domainVerificationDomain.Repo,
) *Service {
	return &Service{
		domainVerificationRepo: domainVerificationRepo,
	}
}
