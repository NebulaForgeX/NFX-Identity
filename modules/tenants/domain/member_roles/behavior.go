package member_roles

import (
	"time"

	"github.com/google/uuid"
)

func (mr *MemberRole) Revoke(revokedBy uuid.UUID, reason string) error {
	if mr.RevokedAt() != nil {
		return ErrMemberRoleAlreadyRevoked
	}

	now := time.Now().UTC()
	mr.state.RevokedAt = &now
	mr.state.RevokedBy = &revokedBy
	mr.state.RevokeReason = &reason
	return nil
}

func (mr *MemberRole) IsExpired() bool {
	if mr.ExpiresAt() == nil {
		return false
	}
	return time.Now().UTC().After(*mr.ExpiresAt())
}

func (mr *MemberRole) IsRevoked() bool {
	return mr.RevokedAt() != nil
}

func (mr *MemberRole) IsActive() bool {
	return !mr.IsExpired() && !mr.IsRevoked()
}
