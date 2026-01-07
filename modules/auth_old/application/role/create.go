package role

import (
	"context"
	roleCommands "nfxid/modules/auth/application/role/commands"
	roleDomain "nfxid/modules/auth/domain/role"
	roleDomainErrors "nfxid/modules/auth/domain/role/errors"
)

func (s *Service) CreateRole(ctx context.Context, cmd roleCommands.CreateRoleCmd) (*roleDomain.Role, error) {
	// 检查角色名是否已存在
	if exists, _ := s.roleRepo.Check.ByName(ctx, cmd.Editable.Name); exists {
		return nil, roleDomainErrors.ErrRoleNameExists
	}

	// 使用 domain factory 创建实体
	r, err := roleDomain.NewRole(roleDomain.NewRoleParams{
		Editable: cmd.Editable,
		IsSystem: cmd.IsSystem,
	})
	if err != nil {
		return nil, err
	}

	if err := s.roleRepo.Create.New(ctx, r); err != nil {
		return nil, err
	}

	return r, nil
}
