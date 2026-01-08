package mfa_factors

import (
	"time"

	"github.com/google/uuid"
)

type NewMFAFactorParams struct {
	FactorID         string
	TenantID         uuid.UUID
	UserID           uuid.UUID
	Type             MFAType
	SecretEncrypted  *string
	Phone            *string
	Email            *string
	Name             *string
	Enabled          bool
	RecoveryCodesHash *string
}

func NewMFAFactor(p NewMFAFactorParams) (*MFAFactor, error) {
	if err := validateMFAFactorParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return NewMFAFactorFromState(MFAFactorState{
		ID:                id,
		FactorID:          p.FactorID,
		TenantID:          p.TenantID,
		UserID:            p.UserID,
		Type:              p.Type,
		SecretEncrypted:   p.SecretEncrypted,
		Phone:             p.Phone,
		Email:             p.Email,
		Name:              p.Name,
		Enabled:           p.Enabled,
		RecoveryCodesHash: p.RecoveryCodesHash,
		CreatedAt:         now,
		UpdatedAt:         now,
	}), nil
}

func NewMFAFactorFromState(st MFAFactorState) *MFAFactor {
	return &MFAFactor{state: st}
}

func validateMFAFactorParams(p NewMFAFactorParams) error {
	if p.FactorID == "" {
		return ErrFactorIDRequired
	}
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if p.TenantID == uuid.Nil {
		return ErrTenantIDRequired
	}
	if p.Type == "" {
		return ErrTypeRequired
	}
	validTypes := map[MFAType]struct{}{
		MFATypeTOTP:       {},
		MFATypeSMS:        {},
		MFATypeEmail:      {},
		MFATypeWebAuthn:   {},
		MFATypeBackupCode: {},
	}
	if _, ok := validTypes[p.Type]; !ok {
		return ErrInvalidMFAType
	}
	return nil
}
