package invitations

import (
	"context"
	invitationCommands "nfxid/modules/tenants/application/invitations/commands"
)

// AcceptInvitation 接受邀请
func (s *Service) AcceptInvitation(ctx context.Context, cmd invitationCommands.AcceptInvitationCmd) error {
	// Get domain entity
	invitation, err := s.invitationRepo.Get.ByInviteID(ctx, cmd.InviteID)
	if err != nil {
		return err
	}

	// Accept domain entity
	if err := invitation.Accept(cmd.UserID); err != nil {
		return err
	}

	// Save to repository
	return s.invitationRepo.Update.Accept(ctx, cmd.InviteID, cmd.UserID)
}

// RevokeInvitation 撤销邀请
func (s *Service) RevokeInvitation(ctx context.Context, cmd invitationCommands.RevokeInvitationCmd) error {
	// Get domain entity
	invitation, err := s.invitationRepo.Get.ByInviteID(ctx, cmd.InviteID)
	if err != nil {
		return err
	}

	// Revoke domain entity
	if err := invitation.Revoke(cmd.RevokedBy, cmd.RevokeReason); err != nil {
		return err
	}

	// Save to repository
	return s.invitationRepo.Update.Revoke(ctx, cmd.InviteID, cmd.RevokedBy, cmd.RevokeReason)
}
