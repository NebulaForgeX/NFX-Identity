package mfa_factors

import "github.com/google/uuid"

func (mf *MFAFactor) Validate() error {
	if mf.FactorID() == "" {
		return ErrFactorIDRequired
	}
	if mf.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	if mf.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	if mf.Type() == "" {
		return ErrTypeRequired
	}
	validTypes := map[MFAType]struct{}{
		MFATypeTOTP:       {},
		MFATypeSMS:        {},
		MFATypeEmail:      {},
		MFATypeWebAuthn:   {},
		MFATypeBackupCode: {},
	}
	if _, ok := validTypes[mf.Type()]; !ok {
		return ErrInvalidMFAType
	}
	return nil
}
