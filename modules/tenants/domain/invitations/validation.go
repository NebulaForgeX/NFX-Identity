package invitations

import "github.com/google/uuid"

func (i *Invitation) Validate() error {
	if i.InviteID() == "" {
		return ErrInviteIDRequired
	}
	if i.TenantID() == uuid.Nil {
		return ErrTenantIDRequired
	}
	if i.Email() == "" {
		return ErrEmailRequired
	}
	if i.TokenHash() == "" {
		return ErrTokenHashRequired
	}
	if i.ExpiresAt().IsZero() {
		return ErrExpiresAtRequired
	}
	if i.InvitedBy() == uuid.Nil {
		return ErrInvitedByRequired
	}
	validStatuses := map[InvitationStatus]struct{}{
		InvitationStatusPending:  {},
		InvitationStatusAccepted: {},
		InvitationStatusExpired:  {},
		InvitationStatusRevoked:  {},
	}
	if _, ok := validStatuses[i.Status()]; !ok {
		return ErrInvalidInvitationStatus
	}
	return nil
}
