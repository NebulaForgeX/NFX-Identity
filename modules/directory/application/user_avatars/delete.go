package user_avatars

import (
	"context"
	userAvatarCommands "nfxidentity/modules/directory/application/user_avatars/commands"
)

// DeleteUserAvatar 删除用户头像
func (s *Service) DeleteUserAvatar(ctx context.Context, cmd userAvatarCommands.DeleteUserAvatarCmd) error {
	return s.userAvatarRepo.Delete.ByUserID(ctx, cmd.UserID)
}
