package member_app_roles

import (
	"context"
	memberAppRoleCommands "nfxid/modules/tenants/application/member_app_roles/commands"
)

// DeleteMemberAppRole 删除成员应用角色
func (s *Service) DeleteMemberAppRole(ctx context.Context, cmd memberAppRoleCommands.DeleteMemberAppRoleCmd) error {
	// Delete from repository (hard delete)
	return s.memberAppRoleRepo.Delete.ByID(ctx, cmd.MemberAppRoleID)
}
