package domain_verifications

import (
	"time"

	"github.com/google/uuid"
)

type NewDomainVerificationParams struct {
	TenantID           uuid.UUID
	Domain             string
	VerificationMethod VerificationMethod
	VerificationToken  *string
	Status             VerificationStatus
	ExpiresAt          *time.Time
	CreatedBy          *uuid.UUID
	Metadata           map[string]interface{}
}

func NewDomainVerification(p NewDomainVerificationParams) (*DomainVerification, error) {
	if err := validateDomainVerificationParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	method := p.VerificationMethod
	if method == "" {
		method = VerificationMethodDNS
	}

	status := p.Status
	if status == "" {
		status = VerificationStatusPending
	}

	now := time.Now().UTC()
	return NewDomainVerificationFromState(DomainVerificationState{
		ID:                 id,
		TenantID:           p.TenantID,
		Domain:             p.Domain,
		VerificationMethod: method,
		VerificationToken:  p.VerificationToken,
		Status:             status,
		ExpiresAt:          p.ExpiresAt,
		CreatedAt:          now,
		CreatedBy:          p.CreatedBy,
		Metadata:           p.Metadata,
	}), nil
}

func NewDomainVerificationFromState(st DomainVerificationState) *DomainVerification {
	return &DomainVerification{state: st}
}

func validateDomainVerificationParams(p NewDomainVerificationParams) error {
	if p.TenantID == uuid.Nil {
		return ErrTenantIDRequired
	}
	if p.Domain == "" {
		return ErrDomainRequired
	}
	if p.VerificationMethod != "" {
		validMethods := map[VerificationMethod]struct{}{
			VerificationMethodDNS:  {},
			VerificationMethodTXT:  {},
			VerificationMethodHTML: {},
			VerificationMethodFILE: {},
		}
		if _, ok := validMethods[p.VerificationMethod]; !ok {
			return ErrInvalidVerificationMethod
		}
	}
	if p.Status != "" {
		validStatuses := map[VerificationStatus]struct{}{
			VerificationStatusPending:  {},
			VerificationStatusVerified: {},
			VerificationStatusFailed:   {},
			VerificationStatusExpired:  {},
		}
		if _, ok := validStatuses[p.Status]; !ok {
			return ErrInvalidVerificationStatus
		}
	}
	return nil
}
