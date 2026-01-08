package member_app_roles

import (
	"time"

	"github.com/google/uuid"
)

func (mar *MemberAppRole) Revoke(revokedBy uuid.UUID, reason string) error {
	if mar.RevokedAt() != nil {
		return ErrMemberAppRoleAlreadyRevoked
	}

	now := time.Now().UTC()
	mar.state.RevokedAt = &now
	mar.state.RevokedBy = &revokedBy
	mar.state.RevokeReason = &reason
	return nil
}

func (mar *MemberAppRole) IsExpired() bool {
	if mar.ExpiresAt() == nil {
		return false
	}
	return time.Now().UTC().After(*mar.ExpiresAt())
}

func (mar *MemberAppRole) IsRevoked() bool {
	return mar.RevokedAt() != nil
}

func (mar *MemberAppRole) IsActive() bool {
	return !mar.IsExpired() && !mar.IsRevoked()
}
