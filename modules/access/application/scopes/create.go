package scopes

import (
	"context"
	scopeCommands "nfxid/modules/access/application/scopes/commands"
	scopeDomain "nfxid/modules/access/domain/scopes"
)

// CreateScope 创建作用域
func (s *Service) CreateScope(ctx context.Context, cmd scopeCommands.CreateScopeCmd) error {
	// Check if scope already exists
	if exists, _ := s.scopeRepo.Check.ByScope(ctx, cmd.Scope); exists {
		return scopeDomain.ErrScopeAlreadyExists
	}

	// Create domain entity
	scope, err := scopeDomain.NewScope(scopeDomain.NewScopeParams{
		Scope:       cmd.Scope,
		Description: cmd.Description,
		IsSystem:    cmd.IsSystem,
	})
	if err != nil {
		return err
	}

	// Save to repository
	return s.scopeRepo.Create.New(ctx, scope)
}
