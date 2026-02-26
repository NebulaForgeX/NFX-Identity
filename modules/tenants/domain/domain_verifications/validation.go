package domain_verifications

import (
	tenantsErr "nfxidentity/errors/src/tenants"

	"github.com/google/uuid"
)

func (dv *DomainVerification) Validate() error {
	if dv.TenantID() == uuid.Nil {
		return tenantsErr.ErrTenantIDRequired
	}
	if dv.Domain() == "" {
		return tenantsErr.ErrDomainRequired
	}
	validMethods := map[VerificationMethod]struct{}{
		VerificationMethodDNS:  {},
		VerificationMethodTXT:  {},
		VerificationMethodHTML: {},
		VerificationMethodFILE: {},
	}
	if _, ok := validMethods[dv.VerificationMethod()]; !ok {
		return tenantsErr.ErrInvalidVerificationMethod
	}
	validStatuses := map[VerificationStatus]struct{}{
		VerificationStatusPending:  {},
		VerificationStatusVerified: {},
		VerificationStatusFailed:   {},
		VerificationStatusExpired:  {},
	}
	if _, ok := validStatuses[dv.Status()]; !ok {
		return tenantsErr.ErrInvalidVerificationStatus
	}
	return nil
}
