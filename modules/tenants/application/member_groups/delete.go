package member_groups

import (
	"context"
	memberGroupCommands "nfxid/modules/tenants/application/member_groups/commands"
)

// DeleteMemberGroup 删除成员组
func (s *Service) DeleteMemberGroup(ctx context.Context, cmd memberGroupCommands.DeleteMemberGroupCmd) error {
	// Delete from repository (hard delete)
	return s.memberGroupRepo.Delete.ByID(ctx, cmd.MemberGroupID)
}
