package scopes

import (
	"context"
	scopeCommands "nfxid/modules/access/application/scopes/commands"
)

// UpdateScope 更新作用域
func (s *Service) UpdateScope(ctx context.Context, cmd scopeCommands.UpdateScopeCmd) error {
	// Get domain entity
	scope, err := s.scopeRepo.Get.ByScope(ctx, cmd.Scope)
	if err != nil {
		return err
	}

	// Update domain entity
	if err := scope.Update(cmd.Description); err != nil {
		return err
	}

	// Save to repository
	return s.scopeRepo.Update.Generic(ctx, scope)
}
