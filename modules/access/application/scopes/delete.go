package scopes

import (
	"context"
	scopeCommands "nfxid/modules/access/application/scopes/commands"
)

// DeleteScope 删除作用域（软删除）
func (s *Service) DeleteScope(ctx context.Context, cmd scopeCommands.DeleteScopeCmd) error {
	// Get domain entity
	scope, err := s.scopeRepo.Get.ByScope(ctx, cmd.Scope)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := scope.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.scopeRepo.Update.Generic(ctx, scope)
}
