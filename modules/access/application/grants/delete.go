package grants

import (
	"context"
	grantCommands "nfxid/modules/access/application/grants/commands"
)

// DeleteGrant 删除授权
func (s *Service) DeleteGrant(ctx context.Context, cmd grantCommands.DeleteGrantCmd) error {
	// Get domain entity
	grant, err := s.grantRepo.Get.ByID(ctx, cmd.GrantID)
	if err != nil {
		return err
	}

	// Delete from repository (hard delete)
	return s.grantRepo.Delete.ByID(ctx, grant.ID())
}

// RevokeGrant 撤销授权
func (s *Service) RevokeGrant(ctx context.Context, cmd grantCommands.RevokeGrantCmd) error {
	// Get domain entity
	grant, err := s.grantRepo.Get.ByID(ctx, cmd.GrantID)
	if err != nil {
		return err
	}

	// Revoke domain entity
	reason := ""
	if cmd.RevokeReason != nil {
		reason = *cmd.RevokeReason
	}
	if err := grant.Revoke(cmd.RevokedBy, reason); err != nil {
		return err
	}

	// Save to repository
	return s.grantRepo.Update.Generic(ctx, grant)
}
