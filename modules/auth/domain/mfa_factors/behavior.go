package mfa_factors

import (
	"time"
)

func (mf *MFAFactor) Enable() error {
	if mf.DeletedAt() != nil {
		return ErrMFAFactorNotFound
	}
	mf.state.Enabled = true
	mf.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (mf *MFAFactor) Disable() error {
	if mf.DeletedAt() != nil {
		return ErrMFAFactorNotFound
	}
	mf.state.Enabled = false
	mf.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (mf *MFAFactor) UpdateLastUsed() error {
	if mf.DeletedAt() != nil {
		return ErrMFAFactorNotFound
	}
	now := time.Now().UTC()
	mf.state.LastUsedAt = &now
	mf.state.UpdatedAt = now
	return nil
}

func (mf *MFAFactor) UpdateName(name *string) error {
	if mf.DeletedAt() != nil {
		return ErrMFAFactorNotFound
	}
	mf.state.Name = name
	mf.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (mf *MFAFactor) UpdateRecoveryCodesHash(hash *string) error {
	if mf.DeletedAt() != nil {
		return ErrMFAFactorNotFound
	}
	mf.state.RecoveryCodesHash = hash
	mf.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (mf *MFAFactor) Delete() error {
	if mf.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	mf.state.DeletedAt = &now
	mf.state.UpdatedAt = now
	return nil
}
