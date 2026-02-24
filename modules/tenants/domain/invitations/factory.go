package invitations

import (
	tenantsErr "nfxid/errors/src/tenants"
	"time"

	"github.com/google/uuid"
)

type NewInvitationParams struct {
	InviteID  string
	TenantID  uuid.UUID
	Email     string
	TokenHash string
	ExpiresAt time.Time
	Status    InvitationStatus
	InvitedBy uuid.UUID
	RoleIDs   []uuid.UUID
	Metadata  map[string]interface{}
}

func NewInvitation(p NewInvitationParams) (*Invitation, error) {
	if err := validateInvitationParams(p); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	status := p.Status
	if status == "" {
		status = InvitationStatusPending
	}

	now := time.Now().UTC()
	return NewInvitationFromState(InvitationState{
		ID:        id,
		InviteID:  p.InviteID,
		TenantID:  p.TenantID,
		Email:     p.Email,
		TokenHash: p.TokenHash,
		ExpiresAt: p.ExpiresAt,
		Status:    status,
		InvitedBy: p.InvitedBy,
		InvitedAt: now,
		RoleIDs:   p.RoleIDs,
		Metadata:  p.Metadata,
	}), nil
}

func NewInvitationFromState(st InvitationState) *Invitation {
	return &Invitation{state: st}
}

func validateInvitationParams(p NewInvitationParams) error {
	if p.InviteID == "" {
		return tenantsErr.ErrInviteIDRequired
	}
	if p.TenantID == uuid.Nil {
		return tenantsErr.ErrTenantIDRequired
	}
	if p.Email == "" {
		return tenantsErr.ErrEmailRequired
	}
	if p.TokenHash == "" {
		return tenantsErr.ErrTokenHashRequired
	}
	if p.ExpiresAt.IsZero() {
		return tenantsErr.ErrExpiresAtRequired
	}
	if p.InvitedBy == uuid.Nil {
		return tenantsErr.ErrInvitedByRequired
	}
	if p.Status != "" {
		validStatuses := map[InvitationStatus]struct{}{
			InvitationStatusPending:  {},
			InvitationStatusAccepted: {},
			InvitationStatusExpired:  {},
			InvitationStatusRevoked:  {},
		}
		if _, ok := validStatuses[p.Status]; !ok {
			return tenantsErr.ErrInvalidInvitationStatus
		}
	}
	return nil
}
