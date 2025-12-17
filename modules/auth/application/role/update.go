package role

import (
	"context"
	roleCommands "nfxid/modules/auth/application/role/commands"
	roleDomainErrors "nfxid/modules/auth/domain/role/errors"
)

func (s *Service) UpdateRole(ctx context.Context, cmd roleCommands.UpdateRoleCmd) error {
	r, err := s.roleRepo.Get.ByID(ctx, cmd.RoleID)
	if err != nil {
		return err
	}

	// 检查新角色名是否已存在（排除当前角色）
	if cmd.Editable.Name != r.Editable().Name {
		if exists, _ := s.roleRepo.Check.ByName(ctx, cmd.Editable.Name); exists {
			return roleDomainErrors.ErrRoleNameExists
		}
	}

	if err := r.EnsureEditable(cmd.Editable); err != nil {
		return err
	}

	if err := r.Update(cmd.Editable); err != nil {
		return err
	}

	if err := s.roleRepo.Update.Generic(ctx, r); err != nil {
		return err
	}

	return nil
}
