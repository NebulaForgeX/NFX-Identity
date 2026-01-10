package client_credentials

import (
	"context"
	clientCredentialCommands "nfxid/modules/clients/application/client_credentials/commands"
)

// DeleteClientCredential 删除客户端凭证
func (s *Service) DeleteClientCredential(ctx context.Context, cmd clientCredentialCommands.DeleteClientCredentialCmd) error {
	// Delete from repository (hard delete by client_id)
	return s.clientCredentialRepo.Delete.ByClientID(ctx, cmd.ClientID)
}
