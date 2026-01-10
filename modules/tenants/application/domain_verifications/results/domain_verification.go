package results

import (
	"time"

	"nfxid/modules/tenants/domain/domain_verifications"

	"github.com/google/uuid"
)

type DomainVerificationRO struct {
	ID                 uuid.UUID
	TenantID           uuid.UUID
	Domain             string
	VerificationMethod domain_verifications.VerificationMethod
	VerificationToken  *string
	Status             domain_verifications.VerificationStatus
	VerifiedAt         *time.Time
	ExpiresAt          *time.Time
	CreatedAt          time.Time
	CreatedBy          *uuid.UUID
	Metadata           map[string]interface{}
}

// DomainVerificationMapper 将 Domain DomainVerification 转换为 Application DomainVerificationRO
func DomainVerificationMapper(dv *domain_verifications.DomainVerification) DomainVerificationRO {
	if dv == nil {
		return DomainVerificationRO{}
	}

	return DomainVerificationRO{
		ID:                 dv.ID(),
		TenantID:           dv.TenantID(),
		Domain:             dv.Domain(),
		VerificationMethod: dv.VerificationMethod(),
		VerificationToken:  dv.VerificationToken(),
		Status:             dv.Status(),
		VerifiedAt:         dv.VerifiedAt(),
		ExpiresAt:          dv.ExpiresAt(),
		CreatedAt:          dv.CreatedAt(),
		CreatedBy:          dv.CreatedBy(),
		Metadata:           dv.Metadata(),
	}
}
