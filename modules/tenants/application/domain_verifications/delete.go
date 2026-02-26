package domain_verifications

import (
	"context"
	domainVerificationCommands "nfxidentity/modules/tenants/application/domain_verifications/commands"
)

// DeleteDomainVerification 删除域名验证
func (s *Service) DeleteDomainVerification(ctx context.Context, cmd domainVerificationCommands.DeleteDomainVerificationCmd) error {
	// Delete from repository (hard delete)
	return s.domainVerificationRepo.Delete.ByID(ctx, cmd.DomainVerificationID)
}
