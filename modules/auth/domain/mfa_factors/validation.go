package mfa_factors

import (
	authErr "nfxidentity/errors/src/auth"

	"github.com/google/uuid"
)

func (mf *MFAFactor) Validate() error {
	if mf.FactorID() == "" {
		return authErr.ErrFactorIDRequired
	}
	if mf.UserID() == uuid.Nil {
		return authErr.ErrUserIDRequired
	}
	if mf.TenantID() == uuid.Nil {
		return authErr.ErrTenantIDRequired
	}
	if mf.Type() == "" {
		return authErr.ErrTypeRequired
	}
	validTypes := map[MFAType]struct{}{
		MFATypeTOTP:       {},
		MFATypeSMS:        {},
		MFATypeEmail:      {},
		MFATypeWebAuthn:   {},
		MFATypeBackupCode: {},
	}
	if _, ok := validTypes[mf.Type()]; !ok {
		return authErr.ErrInvalidMFAType
	}
	return nil
}
