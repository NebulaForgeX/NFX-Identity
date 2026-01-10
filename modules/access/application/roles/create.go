package roles

import (
	"context"
	roleCommands "nfxid/modules/access/application/roles/commands"
	roleDomain "nfxid/modules/access/domain/roles"

	"github.com/google/uuid"
)

// CreateRole 创建角色
func (s *Service) CreateRole(ctx context.Context, cmd roleCommands.CreateRoleCmd) (uuid.UUID, error) {
	// Check if role key already exists
	if exists, _ := s.roleRepo.Check.ByKey(ctx, cmd.Key); exists {
		return uuid.Nil, roleDomain.ErrRoleKeyExists
	}

	// Create domain entity
	role, err := roleDomain.NewRole(roleDomain.NewRoleParams{
		Key:         cmd.Key,
		Name:        cmd.Name,
		Description: cmd.Description,
		ScopeType:   cmd.ScopeType,
		IsSystem:    cmd.IsSystem,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.roleRepo.Create.New(ctx, role); err != nil {
		return uuid.Nil, err
	}

	return role.ID(), nil
}
