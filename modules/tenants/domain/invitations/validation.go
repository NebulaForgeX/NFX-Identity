package invitations

import (
	tenantsErr "nfxid/errors/src/tenants"

	"github.com/google/uuid"
)

func (i *Invitation) Validate() error {
	if i.InviteID() == "" {
		return tenantsErr.ErrInviteIDRequired
	}
	if i.TenantID() == uuid.Nil {
		return tenantsErr.ErrTenantIDRequired
	}
	if i.Email() == "" {
		return tenantsErr.ErrEmailRequired
	}
	if i.TokenHash() == "" {
		return tenantsErr.ErrTokenHashRequired
	}
	if i.ExpiresAt().IsZero() {
		return tenantsErr.ErrExpiresAtRequired
	}
	if i.InvitedBy() == uuid.Nil {
		return tenantsErr.ErrInvitedByRequired
	}
	validStatuses := map[InvitationStatus]struct{}{
		InvitationStatusPending:  {},
		InvitationStatusAccepted: {},
		InvitationStatusExpired:  {},
		InvitationStatusRevoked:  {},
	}
	if _, ok := validStatuses[i.Status()]; !ok {
		return tenantsErr.ErrInvalidInvitationStatus
	}
	return nil
}
