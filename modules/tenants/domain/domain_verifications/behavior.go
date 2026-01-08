package domain_verifications

import (
	"time"
)

func (dv *DomainVerification) Verify() error {
	if dv.Status() == VerificationStatusVerified {
		return nil // already verified
	}
	if dv.Status() == VerificationStatusFailed {
		return ErrDomainVerificationExpired
	}
	if dv.IsExpired() {
		return ErrDomainVerificationExpired
	}

	now := time.Now().UTC()
	dv.state.Status = VerificationStatusVerified
	dv.state.VerifiedAt = &now
	return nil
}

func (dv *DomainVerification) Fail() error {
	if dv.Status() == VerificationStatusVerified {
		return nil // cannot fail a verified domain
	}

	dv.state.Status = VerificationStatusFailed
	return nil
}

func (dv *DomainVerification) UpdateVerificationToken(token *string) error {
	dv.state.VerificationToken = token
	return nil
}

func (dv *DomainVerification) UpdateMetadata(metadata map[string]interface{}) error {
	if metadata == nil {
		return nil
	}
	dv.state.Metadata = metadata
	return nil
}

func (dv *DomainVerification) IsExpired() bool {
	if dv.ExpiresAt() == nil {
		return false
	}
	return time.Now().UTC().After(*dv.ExpiresAt())
}

func (dv *DomainVerification) IsPending() bool {
	return dv.Status() == VerificationStatusPending && !dv.IsExpired()
}
