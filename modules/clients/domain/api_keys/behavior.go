package api_keys

import (
	"time"

	"github.com/google/uuid"
)

func (ak *APIKey) UpdateLastUsed() error {
	if ak.RevokedAt() != nil {
		return ErrAPIKeyAlreadyRevoked
	}
	if ak.IsExpired() {
		return ErrAPIKeyExpired
	}

	now := time.Now().UTC()
	ak.state.LastUsedAt = &now
	return nil
}

func (ak *APIKey) Revoke(revokedBy uuid.UUID, reason string) error {
	if ak.RevokedAt() != nil {
		return ErrAPIKeyAlreadyRevoked
	}

	now := time.Now().UTC()
	ak.state.RevokedAt = &now
	ak.state.RevokedBy = &revokedBy
	ak.state.RevokeReason = &reason
	ak.state.Status = APIKeyStatusRevoked
	return nil
}

func (ak *APIKey) UpdateMetadata(metadata map[string]interface{}) error {
	if metadata == nil {
		return nil
	}
	ak.state.Metadata = metadata
	return nil
}

func (ak *APIKey) IsExpired() bool {
	if ak.ExpiresAt() == nil {
		return false
	}
	return time.Now().UTC().After(*ak.ExpiresAt())
}

func (ak *APIKey) IsRevoked() bool {
	return ak.RevokedAt() != nil
}

func (ak *APIKey) IsValid() bool {
	return ak.Status() == APIKeyStatusActive && !ak.IsExpired() && !ak.IsRevoked()
}
