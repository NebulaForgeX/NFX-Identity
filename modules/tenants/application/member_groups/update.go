package member_groups

import (
	"context"
	memberGroupCommands "nfxid/modules/tenants/application/member_groups/commands"
)

// RevokeMemberGroup 撤销成员组
func (s *Service) RevokeMemberGroup(ctx context.Context, cmd memberGroupCommands.RevokeMemberGroupCmd) error {
	// Get domain entity
	memberGroup, err := s.memberGroupRepo.Get.ByID(ctx, cmd.MemberGroupID)
	if err != nil {
		return err
	}

	// Revoke domain entity
	if err := memberGroup.Revoke(cmd.RevokedBy); err != nil {
		return err
	}

	// Save to repository
	return s.memberGroupRepo.Update.Revoke(ctx, cmd.MemberGroupID, cmd.RevokedBy)
}
