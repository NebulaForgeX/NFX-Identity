package members

import (
	"context"
	memberCommands "nfxid/modules/tenants/application/members/commands"
)

// DeleteMember 删除成员
func (s *Service) DeleteMember(ctx context.Context, cmd memberCommands.DeleteMemberCmd) error {
	// Delete from repository (hard delete)
	return s.memberRepo.Delete.ByID(ctx, cmd.MemberID)
}
