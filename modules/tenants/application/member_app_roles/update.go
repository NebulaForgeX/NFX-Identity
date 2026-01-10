package member_app_roles

import (
	"context"
	memberAppRoleCommands "nfxid/modules/tenants/application/member_app_roles/commands"
)

// RevokeMemberAppRole 撤销成员应用角色
func (s *Service) RevokeMemberAppRole(ctx context.Context, cmd memberAppRoleCommands.RevokeMemberAppRoleCmd) error {
	// Get domain entity
	memberAppRole, err := s.memberAppRoleRepo.Get.ByID(ctx, cmd.MemberAppRoleID)
	if err != nil {
		return err
	}

	// Revoke domain entity
	if err := memberAppRole.Revoke(cmd.RevokedBy, cmd.RevokeReason); err != nil {
		return err
	}

	// Save to repository
	return s.memberAppRoleRepo.Update.Revoke(ctx, cmd.MemberAppRoleID, cmd.RevokedBy, cmd.RevokeReason)
}
