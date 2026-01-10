package scope_permissions

import (
	"context"
	scopePermissionCommands "nfxid/modules/access/application/scope_permissions/commands"
	scopePermissionDomain "nfxid/modules/access/domain/scope_permissions"

	"github.com/google/uuid"
)

// CreateScopePermission 创建作用域权限
func (s *Service) CreateScopePermission(ctx context.Context, cmd scopePermissionCommands.CreateScopePermissionCmd) (uuid.UUID, error) {
	// Check if scope permission already exists
	if exists, _ := s.scopePermissionRepo.Check.ByScopeAndPermissionID(ctx, cmd.Scope, cmd.PermissionID); exists {
		return uuid.Nil, scopePermissionDomain.ErrScopePermissionAlreadyExists
	}

	// Create domain entity
	scopePermission, err := scopePermissionDomain.NewScopePermission(scopePermissionDomain.NewScopePermissionParams{
		Scope:        cmd.Scope,
		PermissionID: cmd.PermissionID,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.scopePermissionRepo.Create.New(ctx, scopePermission); err != nil {
		return uuid.Nil, err
	}

	return scopePermission.ID(), nil
}
