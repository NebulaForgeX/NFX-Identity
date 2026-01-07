package user_role

import (
	"context"
	userRoleCommands "nfxid/modules/auth/application/user_role/commands"
	userRoleDomain "nfxid/modules/auth/domain/user_role"
	userRoleDomainErrors "nfxid/modules/auth/domain/user_role/errors"
)

func (s *Service) CreateUserRole(ctx context.Context, cmd userRoleCommands.CreateUserRoleCmd) (*userRoleDomain.UserRole, error) {
	// 检查关联是否已存在
	if exists, _ := s.userRoleRepo.Check.ByUserAndRole(ctx, cmd.UserID, cmd.RoleID); exists {
		return nil, userRoleDomainErrors.ErrUserRoleAlreadyExists
	}

	// 使用 domain factory 创建实体
	ur, err := userRoleDomain.NewUserRole(userRoleDomain.NewUserRoleParams{
		UserID: cmd.UserID,
		RoleID: cmd.RoleID,
	})
	if err != nil {
		return nil, err
	}

	if err := s.userRoleRepo.Create.New(ctx, ur); err != nil {
		return nil, err
	}

	return ur, nil
}
