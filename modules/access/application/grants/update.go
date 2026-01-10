package grants

import (
	"context"
	"time"
	grantCommands "nfxid/modules/access/application/grants/commands"
)

// UpdateGrant 更新授权
func (s *Service) UpdateGrant(ctx context.Context, cmd grantCommands.UpdateGrantCmd) error {
	// Get domain entity
	grant, err := s.grantRepo.Get.ByID(ctx, cmd.GrantID)
	if err != nil {
		return err
	}

	// Update domain entity
	var expiresAt *time.Time
	if cmd.ExpiresAt != nil && *cmd.ExpiresAt != "" {
		parsed, err := time.Parse(time.RFC3339, *cmd.ExpiresAt)
		if err != nil {
			return err
		}
		expiresAt = &parsed
	}

	if err := grant.UpdateExpiresAt(expiresAt); err != nil {
		return err
	}

	// Save to repository
	return s.grantRepo.Update.Generic(ctx, grant)
}
