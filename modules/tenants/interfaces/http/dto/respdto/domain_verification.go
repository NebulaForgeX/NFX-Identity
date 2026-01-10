package respdto

import (
	"time"

	domainVerificationAppResult "nfxid/modules/tenants/application/domain_verifications/results"

	"github.com/google/uuid"
)

type DomainVerificationDTO struct {
	ID                 uuid.UUID              `json:"id"`
	TenantID           uuid.UUID              `json:"tenant_id"`
	Domain             string                 `json:"domain"`
	VerificationMethod string                 `json:"verification_method"`
	VerificationToken  *string                `json:"verification_token,omitempty"`
	Status             string                 `json:"status"`
	VerifiedAt         *time.Time             `json:"verified_at,omitempty"`
	ExpiresAt          *time.Time             `json:"expires_at,omitempty"`
	CreatedAt          time.Time              `json:"created_at"`
	CreatedBy          *uuid.UUID             `json:"created_by,omitempty"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
}

// DomainVerificationROToDTO converts application DomainVerificationRO to response DTO
func DomainVerificationROToDTO(v *domainVerificationAppResult.DomainVerificationRO) *DomainVerificationDTO {
	if v == nil {
		return nil
	}

	return &DomainVerificationDTO{
		ID:                 v.ID,
		TenantID:           v.TenantID,
		Domain:             v.Domain,
		VerificationMethod: string(v.VerificationMethod),
		VerificationToken:  v.VerificationToken,
		Status:             string(v.Status),
		VerifiedAt:         v.VerifiedAt,
		ExpiresAt:          v.ExpiresAt,
		CreatedAt:          v.CreatedAt,
		CreatedBy:          v.CreatedBy,
		Metadata:           v.Metadata,
	}
}

// DomainVerificationListROToDTO converts list of DomainVerificationRO to DTOs
func DomainVerificationListROToDTO(results []domainVerificationAppResult.DomainVerificationRO) []DomainVerificationDTO {
	dtos := make([]DomainVerificationDTO, len(results))
	for i, v := range results {
		if dto := DomainVerificationROToDTO(&v); dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
