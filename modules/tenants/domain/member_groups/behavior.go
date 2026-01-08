package member_groups

import (
	"time"

	"github.com/google/uuid"
)

func (mg *MemberGroup) Revoke(revokedBy uuid.UUID) error {
	if mg.RevokedAt() != nil {
		return ErrMemberGroupAlreadyRevoked
	}

	now := time.Now().UTC()
	mg.state.RevokedAt = &now
	mg.state.RevokedBy = &revokedBy
	return nil
}

func (mg *MemberGroup) IsRevoked() bool {
	return mg.RevokedAt() != nil
}

func (mg *MemberGroup) IsActive() bool {
	return !mg.IsRevoked()
}
