package invitations

import (
	tenantsErr "nfxidentity/errors/src/tenants"
	"time"

	"github.com/google/uuid"
)

func (i *Invitation) Accept(userID uuid.UUID) error {
	if i.Status() == InvitationStatusAccepted {
		return tenantsErr.ErrInvitationAlreadyAccepted
	}
	if i.Status() == InvitationStatusRevoked {
		return tenantsErr.ErrInvitationAlreadyRevoked
	}
	if i.IsExpired() {
		return tenantsErr.ErrInvitationExpired
	}

	now := time.Now().UTC()
	i.state.Status = InvitationStatusAccepted
	i.state.AcceptedByUserID = &userID
	i.state.AcceptedAt = &now
	return nil
}

func (i *Invitation) Revoke(revokedBy uuid.UUID, reason string) error {
	if i.Status() == InvitationStatusRevoked {
		return tenantsErr.ErrInvitationAlreadyRevoked
	}
	if i.Status() == InvitationStatusAccepted {
		return tenantsErr.ErrInvitationAlreadyAccepted
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
		return tenantsErr.ErrInvitationAlreadyAccepted
	}
	i.state.RoleIDs = roleIDs
	return nil
}
