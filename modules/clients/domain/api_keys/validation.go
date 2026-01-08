package api_keys

import "github.com/google/uuid"

func (ak *APIKey) Validate() error {
	if ak.KeyID() == "" {
		return ErrKeyIDRequired
	}
	if ak.AppID() == uuid.Nil {
		return ErrAppIDRequired
	}
	if ak.KeyHash() == "" {
		return ErrKeyHashRequired
	}
	if ak.HashAlg() == "" {
		return ErrHashAlgRequired
	}
	if ak.Name() == "" {
		return ErrNameRequired
	}
	validStatuses := map[APIKeyStatus]struct{}{
		APIKeyStatusActive:  {},
		APIKeyStatusRevoked: {},
		APIKeyStatusExpired: {},
	}
	if _, ok := validStatuses[ak.Status()]; !ok {
		return ErrInvalidAPIKeyStatus
	}
	return nil
}
