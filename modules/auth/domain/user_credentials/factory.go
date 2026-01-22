package user_credentials

import (
	"time"

	"github.com/google/uuid"
)

type NewUserCredentialParams struct {
	UserID             uuid.UUID
	CredentialType     CredentialType
	PasswordHash       *string
	HashAlg            *string
	HashParams         map[string]interface{}
	Status             CredentialStatus
	MustChangePassword bool
}

func NewUserCredential(p NewUserCredentialParams) (*UserCredential, error) {
	if err := validateUserCredentialParams(p); err != nil {
		return nil, err
	}

	credentialType := p.CredentialType
	if credentialType == "" {
		credentialType = CredentialTypePassword
	}

	status := p.Status
	if status == "" {
		status = CredentialStatusActive
	}

	// id 必须等于 UserID（一对一关系，id 直接引用 directory.users.id）
	now := time.Now().UTC()
	var passwordUpdatedAt *time.Time
	if p.PasswordHash != nil {
		passwordUpdatedAt = &now
	}

	return NewUserCredentialFromState(UserCredentialState{
		ID:                 p.UserID, // id 直接引用 directory.users.id
		UserID:             p.UserID,
		CredentialType:     credentialType,
		PasswordHash:       p.PasswordHash,
		HashAlg:            p.HashAlg,
		HashParams:         p.HashParams,
		PasswordUpdatedAt:  passwordUpdatedAt,
		Status:             status,
		MustChangePassword: p.MustChangePassword,
		Version:            1,
		CreatedAt:          now,
		UpdatedAt:          now,
	}), nil
}

func NewUserCredentialFromState(st UserCredentialState) *UserCredential {
	return &UserCredential{state: st}
}

func validateUserCredentialParams(p NewUserCredentialParams) error {
	if p.UserID == uuid.Nil {
		return ErrUserIDRequired
	}
	if p.CredentialType != "" {
		validTypes := map[CredentialType]struct{}{
			CredentialTypePassword:  {},
			CredentialTypePasskey:   {},
			CredentialTypeOauthLink: {},
			CredentialTypeSaml:      {},
			CredentialTypeLdap:      {},
		}
		if _, ok := validTypes[p.CredentialType]; !ok {
			return ErrInvalidCredentialType
		}
	}
	if p.Status != "" {
		validStatuses := map[CredentialStatus]struct{}{
			CredentialStatusActive:   {},
			CredentialStatusDisabled: {},
			CredentialStatusExpired:  {},
		}
		if _, ok := validStatuses[p.Status]; !ok {
			return ErrInvalidCredentialStatus
		}
	}
	return nil
}
