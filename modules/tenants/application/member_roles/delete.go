package member_roles

import (
	"context"
	memberRoleCommands "nfxid/modules/tenants/application/member_roles/commands"
)

// DeleteMemberRole 删除成员角色
func (s *Service) DeleteMemberRole(ctx context.Context, cmd memberRoleCommands.DeleteMemberRoleCmd) error {
	// Delete from repository (hard delete)
	return s.memberRoleRepo.Delete.ByID(ctx, cmd.MemberRoleID)
}
