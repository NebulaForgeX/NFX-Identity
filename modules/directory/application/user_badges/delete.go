package user_badges

import (
	"context"
	userBadgeCommands "nfxid/modules/directory/application/user_badges/commands"
)

// DeleteUserBadge 删除用户徽章
func (s *Service) DeleteUserBadge(ctx context.Context, cmd userBadgeCommands.DeleteUserBadgeCmd) error {
	// Delete from repository (hard delete)
	return s.userBadgeRepo.Delete.ByID(ctx, cmd.UserBadgeID)
}
