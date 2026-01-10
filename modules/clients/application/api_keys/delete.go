package api_keys

import (
	"context"
	apiKeyCommands "nfxid/modules/clients/application/api_keys/commands"
)

// DeleteAPIKey 删除API密钥
func (s *Service) DeleteAPIKey(ctx context.Context, cmd apiKeyCommands.DeleteAPIKeyCmd) error {
	// Delete from repository (hard delete by key_id)
	return s.apiKeyRepo.Delete.ByKeyID(ctx, cmd.KeyID)
}
