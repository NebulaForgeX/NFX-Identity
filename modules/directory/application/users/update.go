package users

import (
	"context"
	userCommands "nfxid/modules/directory/application/users/commands"
)

// UpdateUserStatus 更新用户状态
func (s *Service) UpdateUserStatus(ctx context.Context, cmd userCommands.UpdateUserStatusCmd) error {
	// Get domain entity
	user, err := s.userRepo.Get.ByID(ctx, cmd.UserID)
	if err != nil {
		return err
	}

	// Update status domain entity
	if err := user.UpdateStatus(cmd.Status); err != nil {
		return err
	}

	// Save to repository
	return s.userRepo.Update.Status(ctx, cmd.UserID, cmd.Status)
}

// UpdateUsername 更新用户名
func (s *Service) UpdateUsername(ctx context.Context, cmd userCommands.UpdateUsernameCmd) error {
	// Get domain entity
	user, err := s.userRepo.Get.ByID(ctx, cmd.UserID)
	if err != nil {
		return err
	}

	// Update username domain entity
	if err := user.UpdateUsername(cmd.Username); err != nil {
		return err
	}

	// Save to repository
	return s.userRepo.Update.Generic(ctx, user)
}

// VerifyUser 验证用户
func (s *Service) VerifyUser(ctx context.Context, cmd userCommands.VerifyUserCmd) error {
	// Get domain entity
	user, err := s.userRepo.Get.ByID(ctx, cmd.UserID)
	if err != nil {
		return err
	}

	// Verify domain entity
	if err := user.Verify(); err != nil {
		return err
	}

	// Save to repository
	return s.userRepo.Update.Verify(ctx, cmd.UserID)
}
