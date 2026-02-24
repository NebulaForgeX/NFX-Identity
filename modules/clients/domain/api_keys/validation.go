package api_keys

import (
	clientsErr "nfxid/errors/src/clients"

	"github.com/google/uuid"
)

func (ak *APIKey) Validate() error {
	if ak.KeyID() == "" {
		return clientsErr.ErrKeyIDRequired
	}
	if ak.AppID() == uuid.Nil {
		return clientsErr.ErrAppIDRequired
	}
	if ak.KeyHash() == "" {
		return clientsErr.ErrKeyHashRequired
	}
	if ak.HashAlg() == "" {
		return clientsErr.ErrHashAlgRequired
	}
	if ak.Name() == "" {
		return clientsErr.ErrNameRequired
	}
	validStatuses := map[APIKeyStatus]struct{}{
		APIKeyStatusActive:  {},
		APIKeyStatusRevoked: {},
		APIKeyStatusExpired: {},
	}
	if _, ok := validStatuses[ak.Status()]; !ok {
		return clientsErr.ErrInvalidAPIKeyStatus
	}
	return nil
}
