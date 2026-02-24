package client_credentials

import (
	clientsErr "nfxid/errors/src/clients"

	"github.com/google/uuid"
)

func (cc *ClientCredential) Validate() error {
	if cc.ClientID() == "" {
		return clientsErr.ErrClientIDRequired
	}
	if cc.AppID() == uuid.Nil {
		return clientsErr.ErrAppIDRequired
	}
	if cc.SecretHash() == "" {
		return clientsErr.ErrSecretHashRequired
	}
	if cc.HashAlg() == "" {
		return clientsErr.ErrHashAlgRequired
	}
	validStatuses := map[CredentialStatus]struct{}{
		CredentialStatusActive:   {},
		CredentialStatusExpired:  {},
		CredentialStatusRevoked:  {},
		CredentialStatusRotating: {},
	}
	if _, ok := validStatuses[cc.Status()]; !ok {
		return clientsErr.ErrInvalidCredentialStatus
	}
	return nil
}
