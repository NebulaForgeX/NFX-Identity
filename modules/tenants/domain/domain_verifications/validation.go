package domain_verifications

import "github.com/google/uuid"

func (dv *DomainVerification) Validate() error {
	if dv.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	if dv.Domain() == "" {
		return ErrDomainRequired
	}
	validMethods := map[VerificationMethod]struct{}{
		VerificationMethodDNS:  {},
		VerificationMethodTXT:  {},
		VerificationMethodHTML: {},
		VerificationMethodFILE: {},
	}
	if _, ok := validMethods[dv.VerificationMethod()]; !ok {
		return ErrInvalidVerificationMethod
	}
	validStatuses := map[VerificationStatus]struct{}{
		VerificationStatusPending:  {},
		VerificationStatusVerified: {},
		VerificationStatusFailed:   {},
		VerificationStatusExpired:  {},
	}
	if _, ok := validStatuses[dv.Status()]; !ok {
		return ErrInvalidVerificationStatus
	}
	return nil
}
