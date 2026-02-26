package client_scopes

import (
	"context"
	clientScopeCommands "nfxidentity/modules/clients/application/client_scopes/commands"
)

// DeleteClientScope 删除客户端作用域
func (s *Service) DeleteClientScope(ctx context.Context, cmd clientScopeCommands.DeleteClientScopeCmd) error {
	// Delete from repository (hard delete by id)
	return s.clientScopeRepo.Delete.ByID(ctx, cmd.ClientScopeID)
}
