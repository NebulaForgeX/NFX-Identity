package user_credentials

import (
	"time"
)

func (uc *UserCredential) UpdatePassword(passwordHash string, hashAlg string, hashParams map[string]interface{}) error {
	if uc.DeletedAt() != nil {
		return ErrUserCredentialNotFound
	}
	if passwordHash == "" {
		return ErrPasswordHashRequired
	}

	now := time.Now().UTC()
	uc.state.PasswordHash = &passwordHash
	uc.state.HashAlg = &hashAlg
	uc.state.HashParams = hashParams
	uc.state.PasswordUpdatedAt = &now
	uc.state.MustChangePassword = false
	uc.state.Version++
	uc.state.UpdatedAt = now
	return nil
}

func (uc *UserCredential) UpdateStatus(status CredentialStatus) error {
	if uc.DeletedAt() != nil {
		return ErrUserCredentialNotFound
	}
	validStatuses := map[CredentialStatus]struct{}{
		CredentialStatusActive:   {},
		CredentialStatusDisabled: {},
		CredentialStatusExpired:  {},
	}
	if _, ok := validStatuses[status]; !ok {
		return ErrInvalidCredentialStatus
	}

	uc.state.Status = status
	uc.state.Version++
	uc.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (uc *UserCredential) UpdateLastSuccessLogin() error {
	if uc.DeletedAt() != nil {
		return ErrUserCredentialNotFound
	}

	now := time.Now().UTC()
	uc.state.LastSuccessLoginAt = &now
	uc.state.UpdatedAt = now
	return nil
}

func (uc *UserCredential) SetMustChangePassword(mustChange bool) error {
	if uc.DeletedAt() != nil {
		return ErrUserCredentialNotFound
	}

	uc.state.MustChangePassword = mustChange
	uc.state.Version++
	uc.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (uc *UserCredential) Delete() error {
	if uc.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	uc.state.DeletedAt = &now
	uc.state.UpdatedAt = now
	return nil
}
