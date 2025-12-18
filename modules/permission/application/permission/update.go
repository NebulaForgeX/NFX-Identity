package permission

import (
	"context"
	permissionCommands "nfxid/modules/permission/application/permission/commands"
	permissionDomain "nfxid/modules/permission/domain/permission"
	permissionDomainErrors "nfxid/modules/permission/domain/permission/errors"
)

func (s *Service) UpdatePermission(ctx context.Context, cmd permissionCommands.UpdatePermissionCmd) error {
	p, err := s.permissionRepo.Get.ByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	if p == nil {
		return permissionDomainErrors.ErrPermissionNotFound
	}

	if err := p.Update(permissionDomain.PermissionEditable{
		Tag:         cmd.Tag,
		Name:        cmd.Name,
		Description: cmd.Description,
		Category:    cmd.Category,
	}); err != nil {
		return err
	}

	return s.permissionRepo.Update.Generic(ctx, p)
}
