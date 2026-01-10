package domain_verifications

import (
	"context"
	domainVerificationCommands "nfxid/modules/tenants/application/domain_verifications/commands"
)

// VerifyDomain 验证域名
func (s *Service) VerifyDomain(ctx context.Context, cmd domainVerificationCommands.VerifyDomainCmd) error {
	// Get domain entity
	domainVerification, err := s.domainVerificationRepo.Get.ByID(ctx, cmd.DomainVerificationID)
	if err != nil {
		return err
	}

	// Verify domain entity
	if err := domainVerification.Verify(); err != nil {
		return err
	}

	// Save to repository
	return s.domainVerificationRepo.Update.Verify(ctx, cmd.DomainVerificationID)
}

// FailDomainVerification 标记域名验证失败
func (s *Service) FailDomainVerification(ctx context.Context, cmd domainVerificationCommands.FailDomainVerificationCmd) error {
	// Get domain entity
	domainVerification, err := s.domainVerificationRepo.Get.ByID(ctx, cmd.DomainVerificationID)
	if err != nil {
		return err
	}

	// Fail domain entity
	if err := domainVerification.Fail(); err != nil {
		return err
	}

	// Save to repository
	return s.domainVerificationRepo.Update.Fail(ctx, cmd.DomainVerificationID)
}
