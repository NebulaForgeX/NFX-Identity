package permission

import (
	"context"
	permissionCommands "nfxid/modules/permission/application/permission/commands"
	permissionDomainErrors "nfxid/modules/permission/domain/permission/errors"
)

func (s *Service) DeletePermission(ctx context.Context, cmd permissionCommands.DeletePermissionCmd) error {
	p, err := s.permissionRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	if p == nil {
		return permissionDomainErrors.ErrPermissionNotFound
	}

	if err := p.Delete(); err != nil {
		return err
	}

	return s.permissionRepo.Update(ctx, p)
}

