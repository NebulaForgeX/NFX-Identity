package reqdto

import (
	invitationAppCommands "nfxid/modules/tenants/application/invitations/commands"
	invitationDomain "nfxid/modules/tenants/domain/invitations"

	"github.com/google/uuid"
)

type InvitationCreateRequestDTO struct {
	InviteID  string                 `json:"invite_id" validate:"required"`
	TenantID  uuid.UUID              `json:"tenant_id" validate:"required,uuid"`
	Email     string                 `json:"email" validate:"required,email"`
	TokenHash string                 `json:"token_hash" validate:"required"`
	ExpiresAt string                 `json:"expires_at" validate:"required"`
	Status    string                 `json:"status,omitempty"`
	InvitedBy uuid.UUID              `json:"invited_by" validate:"required,uuid"`
	RoleIDs   []uuid.UUID            `json:"role_ids,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type InvitationAcceptRequestDTO struct {
	InviteID string    `uri:"invite_id" validate:"required"`
	UserID   uuid.UUID `json:"user_id" validate:"required,uuid"`
}

type InvitationRevokeRequestDTO struct {
	InviteID     string    `uri:"invite_id" validate:"required"`
	RevokedBy    uuid.UUID `json:"revoked_by" validate:"required,uuid"`
	RevokeReason string    `json:"revoke_reason,omitempty"`
}

type InvitationByIDRequestDTO struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type InvitationByInviteIDRequestDTO struct {
	InviteID string `uri:"invite_id" validate:"required"`
}

func (r *InvitationCreateRequestDTO) ToCreateCmd() invitationAppCommands.CreateInvitationCmd {
	cmd := invitationAppCommands.CreateInvitationCmd{
		InviteID:  r.InviteID,
		TenantID:  r.TenantID,
		Email:     r.Email,
		TokenHash: r.TokenHash,
		ExpiresAt: r.ExpiresAt,
		InvitedBy: r.InvitedBy,
		RoleIDs:   r.RoleIDs,
		Metadata:  r.Metadata,
	}

	// Parse status
	if r.Status != "" {
		cmd.Status = invitationDomain.InvitationStatus(r.Status)
	} else {
		cmd.Status = invitationDomain.InvitationStatusPending
	}

	return cmd
}

func (r *InvitationAcceptRequestDTO) ToAcceptCmd() invitationAppCommands.AcceptInvitationCmd {
	return invitationAppCommands.AcceptInvitationCmd{
		InviteID: r.InviteID,
		UserID:   r.UserID,
	}
}

func (r *InvitationRevokeRequestDTO) ToRevokeCmd() invitationAppCommands.RevokeInvitationCmd {
	return invitationAppCommands.RevokeInvitationCmd{
		InviteID:     r.InviteID,
		RevokedBy:    r.RevokedBy,
		RevokeReason: r.RevokeReason,
	}
}
