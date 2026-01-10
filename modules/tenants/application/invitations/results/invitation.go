package results

import (
	"time"

	"nfxid/modules/tenants/domain/invitations"

	"github.com/google/uuid"
)

type InvitationRO struct {
	ID              uuid.UUID
	InviteID        string
	TenantID        uuid.UUID
	Email           string
	TokenHash       string
	ExpiresAt       time.Time
	Status          invitations.InvitationStatus
	InvitedBy       uuid.UUID
	InvitedAt       time.Time
	AcceptedByUserID *uuid.UUID
	AcceptedAt     *time.Time
	RevokedBy       *uuid.UUID
	RevokedAt       *time.Time
	RevokeReason    *string
	RoleIDs         []uuid.UUID
	Metadata        map[string]interface{}
}

// InvitationMapper 将 Domain Invitation 转换为 Application InvitationRO
func InvitationMapper(i *invitations.Invitation) InvitationRO {
	if i == nil {
		return InvitationRO{}
	}

	return InvitationRO{
		ID:              i.ID(),
		InviteID:        i.InviteID(),
		TenantID:        i.TenantID(),
		Email:           i.Email(),
		TokenHash:       i.TokenHash(),
		ExpiresAt:       i.ExpiresAt(),
		Status:          i.Status(),
		InvitedBy:       i.InvitedBy(),
		InvitedAt:       i.InvitedAt(),
		AcceptedByUserID: i.AcceptedByUserID(),
		AcceptedAt:      i.AcceptedAt(),
		RevokedBy:       i.RevokedBy(),
		RevokedAt:       i.RevokedAt(),
		RevokeReason:    i.RevokeReason(),
		RoleIDs:         i.RoleIDs(),
		Metadata:        i.Metadata(),
	}
}
