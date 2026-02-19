package reqdto

import (
	domainVerificationAppCommands "nfxid/modules/tenants/application/domain_verifications/commands"
	domainVerificationDomain "nfxid/modules/tenants/domain/domain_verifications"

	"github.com/google/uuid"
)

type DomainVerificationCreateRequestDTO struct {
	TenantID           uuid.UUID              `json:"tenant_id" validate:"required,uuid"`
	Domain             string                 `json:"domain" validate:"required"`
	VerificationMethod string                 `json:"verification_method" validate:"required"`
	VerificationToken  *string                `json:"verification_token,omitempty"`
	Status             string                 `json:"status,omitempty"`
	ExpiresAt          *string                `json:"expires_at,omitempty"`
	CreatedBy          *uuid.UUID             `json:"created_by,omitempty"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
}

type DomainVerificationVerifyRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type DomainVerificationFailRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

func (r *DomainVerificationCreateRequestDTO) ToCreateCmd() domainVerificationAppCommands.CreateDomainVerificationCmd {
	cmd := domainVerificationAppCommands.CreateDomainVerificationCmd{
		TenantID:          r.TenantID,
		Domain:            r.Domain,
		VerificationToken: r.VerificationToken,
		ExpiresAt:         r.ExpiresAt,
		CreatedBy:         r.CreatedBy,
		Metadata:          r.Metadata,
	}

	// Parse verification method
	if r.VerificationMethod != "" {
		cmd.VerificationMethod = domainVerificationDomain.VerificationMethod(r.VerificationMethod)
	} else {
		cmd.VerificationMethod = domainVerificationDomain.VerificationMethodDNS
	}

	// Parse status
	if r.Status != "" {
		cmd.Status = domainVerificationDomain.VerificationStatus(r.Status)
	} else {
		cmd.Status = domainVerificationDomain.VerificationStatusPending
	}

	return cmd
}

func (r *DomainVerificationVerifyRequestDTO) ToVerifyCmd() domainVerificationAppCommands.VerifyDomainCmd {
	return domainVerificationAppCommands.VerifyDomainCmd{
		DomainVerificationID: r.ID,
	}
}

func (r *DomainVerificationFailRequestDTO) ToFailCmd() domainVerificationAppCommands.FailDomainVerificationCmd {
	return domainVerificationAppCommands.FailDomainVerificationCmd{
		DomainVerificationID: r.ID,
	}
}
