package client_credentials

import "github.com/google/uuid"

func (cc *ClientCredential) Validate() error {
	if cc.ClientID() == "" {
		return ErrClientIDRequired
	}
	if cc.AppID() == uuid.Nil {
		return ErrAppIDRequired
	}
	if cc.SecretHash() == "" {
		return ErrSecretHashRequired
	}
	if cc.HashAlg() == "" {
		return ErrHashAlgRequired
	}
	validStatuses := map[CredentialStatus]struct{}{
		CredentialStatusActive:   {},
		CredentialStatusExpired:  {},
		CredentialStatusRevoked:  {},
		CredentialStatusRotating: {},
	}
	if _, ok := validStatuses[cc.Status()]; !ok {
		return ErrInvalidCredentialStatus
	}
	return nil
}
