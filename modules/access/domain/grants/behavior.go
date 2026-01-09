package grants

import (
	"time"

	"github.com/google/uuid"
)

func (g *Grant) Revoke(revokedBy uuid.UUID, reason string) error {
	if g.RevokedAt() != nil {
		return ErrGrantAlreadyRevoked
	}

	now := time.Now().UTC()
	g.state.RevokedAt = &now
	g.state.RevokedBy = &revokedBy
	if reason != "" {
		g.state.RevokeReason = &reason
	}
	return nil
}

func (g *Grant) UpdateExpiresAt(expiresAt *time.Time) error {
	if g.RevokedAt() != nil {
		return ErrGrantAlreadyRevoked
	}

	g.state.ExpiresAt = expiresAt
	return nil
}
