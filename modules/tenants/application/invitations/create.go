package invitations

import (
	"context"
	"time"
	invitationCommands "nfxid/modules/tenants/application/invitations/commands"
	invitationDomain "nfxid/modules/tenants/domain/invitations"

	"github.com/google/uuid"
)

// CreateInvitation 创建邀请
func (s *Service) CreateInvitation(ctx context.Context, cmd invitationCommands.CreateInvitationCmd) (uuid.UUID, error) {
	// Parse expires at
	expiresAt, err := time.Parse(time.RFC3339, cmd.ExpiresAt)
	if err != nil {
		return uuid.Nil, err
	}

	// Create domain entity
	invitation, err := invitationDomain.NewInvitation(invitationDomain.NewInvitationParams{
		InviteID:  cmd.InviteID,
		TenantID:  cmd.TenantID,
		Email:     cmd.Email,
		TokenHash: cmd.TokenHash,
		ExpiresAt: expiresAt,
		Status:    cmd.Status,
		InvitedBy: cmd.InvitedBy,
		RoleIDs:   cmd.RoleIDs,
		Metadata:  cmd.Metadata,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.invitationRepo.Create.New(ctx, invitation); err != nil {
		return uuid.Nil, err
	}

	return invitation.ID(), nil
}
