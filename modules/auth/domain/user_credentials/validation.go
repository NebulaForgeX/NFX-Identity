package user_credentials

import "github.com/google/uuid"

func (uc *UserCredential) Validate() error {
	if uc.UserID() == uuid.Nil {
		return ErrUserIDRequired
	}
	validTypes := map[CredentialType]struct{}{
		CredentialTypePassword:  {},
		CredentialTypePasskey:   {},
		CredentialTypeOauthLink: {},
		CredentialTypeSaml:      {},
		CredentialTypeLdap:      {},
	}
	if _, ok := validTypes[uc.CredentialType()]; !ok {
		return ErrInvalidCredentialType
	}
	validStatuses := map[CredentialStatus]struct{}{
		CredentialStatusActive:   {},
		CredentialStatusDisabled: {},
		CredentialStatusExpired:  {},
	}
	if _, ok := validStatuses[uc.Status()]; !ok {
		return ErrInvalidCredentialStatus
	}
	return nil
}
