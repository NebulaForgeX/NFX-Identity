package permission

import (
	"context"
	permissionCommands "nfxid/modules/permission/application/permission/commands"
	permissionDomain "nfxid/modules/permission/domain/permission"
)

func (s *Service) CreatePermission(ctx context.Context, cmd permissionCommands.CreatePermissionCmd) (*permissionDomain.Permission, error) {
	p, err := permissionDomain.NewPermission(permissionDomain.NewPermissionParams{
		Editable: permissionDomain.PermissionEditable{
			Tag:         cmd.Tag,
			Name:        cmd.Name,
			Description: cmd.Description,
			Category:    cmd.Category,
		},
		IsSystem: cmd.IsSystem,
	})
	if err != nil {
		return nil, err
	}

	if err := s.permissionRepo.Create(ctx, p); err != nil {
		return nil, err
	}

	return p, nil
}

