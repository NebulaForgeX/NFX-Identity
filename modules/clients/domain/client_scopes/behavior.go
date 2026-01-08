package client_scopes

import (
	"time"

	"github.com/google/uuid"
)

func (cs *ClientScope) Revoke(revokedBy uuid.UUID, reason string) error {
	if cs.RevokedAt() != nil {
		return ErrClientScopeAlreadyRevoked
	}

	now := time.Now().UTC()
	cs.state.RevokedAt = &now
	cs.state.RevokedBy = &revokedBy
	cs.state.RevokeReason = &reason
	return nil
}

func (cs *ClientScope) IsExpired() bool {
	if cs.ExpiresAt() == nil {
		return false
	}
	return time.Now().UTC().After(*cs.ExpiresAt())
}

func (cs *ClientScope) IsRevoked() bool {
	return cs.RevokedAt() != nil
}

func (cs *ClientScope) IsValid() bool {
	return !cs.IsExpired() && !cs.IsRevoked()
}
