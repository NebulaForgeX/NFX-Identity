package domain_verifications

import (
	"context"
	"time"
	domainVerificationCommands "nfxid/modules/tenants/application/domain_verifications/commands"
	domainVerificationDomain "nfxid/modules/tenants/domain/domain_verifications"

	"github.com/google/uuid"
)

// CreateDomainVerification 创建域名验证
func (s *Service) CreateDomainVerification(ctx context.Context, cmd domainVerificationCommands.CreateDomainVerificationCmd) (uuid.UUID, error) {
	var expiresAt *time.Time
	if cmd.ExpiresAt != nil && *cmd.ExpiresAt != "" {
		parsed, err := time.Parse(time.RFC3339, *cmd.ExpiresAt)
		if err != nil {
			return uuid.Nil, err
		}
		expiresAt = &parsed
	}

	// Create domain entity
	domainVerification, err := domainVerificationDomain.NewDomainVerification(domainVerificationDomain.NewDomainVerificationParams{
		TenantID:           cmd.TenantID,
		Domain:             cmd.Domain,
		VerificationMethod: cmd.VerificationMethod,
		VerificationToken:  cmd.VerificationToken,
		Status:             cmd.Status,
		ExpiresAt:          expiresAt,
		CreatedBy:          cmd.CreatedBy,
		Metadata:           cmd.Metadata,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.domainVerificationRepo.Create.New(ctx, domainVerification); err != nil {
		return uuid.Nil, err
	}

	return domainVerification.ID(), nil
}
