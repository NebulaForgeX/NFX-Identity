package user_credentials

import (
	authErr "nfxid/errors/src/auth"

	"github.com/google/uuid"
)

func (uc *UserCredential) Validate() error {
	if uc.UserID() == uuid.Nil {
		return authErr.ErrUserIDRequired
	}
	validTypes := map[CredentialType]struct{}{
		CredentialTypePassword:  {},
		CredentialTypePasskey:   {},
		CredentialTypeOauthLink: {},
		CredentialTypeSaml:      {},
		CredentialTypeLdap:      {},
	}
	if _, ok := validTypes[uc.CredentialType()]; !ok {
		return authErr.ErrInvalidCredentialType
	}
	validStatuses := map[CredentialStatus]struct{}{
		CredentialStatusActive:   {},
		CredentialStatusDisabled: {},
		CredentialStatusExpired:  {},
	}
	if _, ok := validStatuses[uc.Status()]; !ok {
		return authErr.ErrInvalidCredentialStatus
	}
	return nil
}
