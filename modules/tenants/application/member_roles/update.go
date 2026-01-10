package member_roles

import (
	"context"
	memberRoleCommands "nfxid/modules/tenants/application/member_roles/commands"
)

// RevokeMemberRole 撤销成员角色
func (s *Service) RevokeMemberRole(ctx context.Context, cmd memberRoleCommands.RevokeMemberRoleCmd) error {
	// Get domain entity
	memberRole, err := s.memberRoleRepo.Get.ByID(ctx, cmd.MemberRoleID)
	if err != nil {
		return err
	}

	// Revoke domain entity
	if err := memberRole.Revoke(cmd.RevokedBy, cmd.RevokeReason); err != nil {
		return err
	}

	// Save to repository
	return s.memberRoleRepo.Update.Revoke(ctx, cmd.MemberRoleID, cmd.RevokedBy, cmd.RevokeReason)
}
