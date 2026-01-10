package permissions

import (
	"context"
	permissionCommands "nfxid/modules/access/application/permissions/commands"
	permissionDomain "nfxid/modules/access/domain/permissions"

	"github.com/google/uuid"
)

// CreatePermission 创建权限
func (s *Service) CreatePermission(ctx context.Context, cmd permissionCommands.CreatePermissionCmd) (uuid.UUID, error) {
	// Check if permission key already exists
	if exists, _ := s.permissionRepo.Check.ByKey(ctx, cmd.Key); exists {
		return uuid.Nil, permissionDomain.ErrPermissionKeyExists
	}

	// Create domain entity
	permission, err := permissionDomain.NewPermission(permissionDomain.NewPermissionParams{
		Key:         cmd.Key,
		Name:        cmd.Name,
		Description: cmd.Description,
		IsSystem:    cmd.IsSystem,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.permissionRepo.Create.New(ctx, permission); err != nil {
		return uuid.Nil, err
	}

	return permission.ID(), nil
}
