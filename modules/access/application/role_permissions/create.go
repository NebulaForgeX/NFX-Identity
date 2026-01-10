package role_permissions

import (
	"context"
	rolePermissionCommands "nfxid/modules/access/application/role_permissions/commands"
	rolePermissionDomain "nfxid/modules/access/domain/role_permissions"

	"github.com/google/uuid"
)

// CreateRolePermission 创建角色权限
func (s *Service) CreateRolePermission(ctx context.Context, cmd rolePermissionCommands.CreateRolePermissionCmd) (uuid.UUID, error) {
	// Check if role permission already exists
	if exists, _ := s.rolePermissionRepo.Check.ByRoleIDAndPermissionID(ctx, cmd.RoleID, cmd.PermissionID); exists {
		return uuid.Nil, rolePermissionDomain.ErrRolePermissionAlreadyExists
	}

	// Create domain entity
	rolePermission, err := rolePermissionDomain.NewRolePermission(rolePermissionDomain.NewRolePermissionParams{
		RoleID:       cmd.RoleID,
		PermissionID: cmd.PermissionID,
		CreatedBy:    cmd.CreatedBy,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.rolePermissionRepo.Create.New(ctx, rolePermission); err != nil {
		return uuid.Nil, err
	}

	return rolePermission.ID(), nil
}
