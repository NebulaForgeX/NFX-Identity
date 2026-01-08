package invitations

import (
	"time"

	"github.com/google/uuid"
)

func (i *Invitation) Accept(userID uuid.UUID) error {
	if i.Status() == InvitationStatusAccepted {
		return ErrInvitationAlreadyAccepted
	}
	if i.Status() == InvitationStatusRevoked {
		return ErrInvitationAlreadyRevoked
	}
	if i.IsExpired() {
		return ErrInvitationExpired
	}

	now := time.Now().UTC()
	i.state.Status = InvitationStatusAccepted
	i.state.AcceptedByUserID = &userID
	i.state.AcceptedAt = &now
	return nil
}

func (i *Invitation) Revoke(revokedBy uuid.UUID, reason string) error {
	if i.Status() == InvitationStatusRevoked {
		return ErrInvitationAlreadyRevoked
	}
	if i.Status() == InvitationStatusAccepted {
		return ErrInvitationAlreadyAccepted
	}

	now := time.Now().UTC()
	i.state.Status = InvitationStatusRevoked
	i.state.RevokedBy = &revokedBy
	i.state.RevokedAt = &now
	i.state.RevokeReason = &reason
	return nil
}

func (i *Invitation) IsExpired() bool {
	return time.Now().UTC().After(i.ExpiresAt())
}

func (i *Invitation) IsValid() bool {
	return i.Status() == InvitationStatusPending && !i.IsExpired()
}

func (i *Invitation) UpdateRoleIDs(roleIDs []uuid.UUID) error {
	if i.Status() != InvitationStatusPending {
		return ErrInvitationAlreadyAccepted
	}
	i.state.RoleIDs = roleIDs
	return nil
}
