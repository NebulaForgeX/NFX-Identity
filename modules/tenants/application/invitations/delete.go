package invitations

import (
	"context"
	invitationCommands "nfxid/modules/tenants/application/invitations/commands"
)

// DeleteInvitation 删除邀请
func (s *Service) DeleteInvitation(ctx context.Context, cmd invitationCommands.DeleteInvitationCmd) error {
	// Delete from repository (hard delete)
	return s.invitationRepo.Delete.ByInviteID(ctx, cmd.InviteID)
}
