package client_scopes

import (
	"context"
	"time"
	clientScopeCommands "nfxid/modules/clients/application/client_scopes/commands"
	clientScopeDomain "nfxid/modules/clients/domain/client_scopes"

	"github.com/google/uuid"
)

// CreateClientScope 创建客户端作用域
func (s *Service) CreateClientScope(ctx context.Context, cmd clientScopeCommands.CreateClientScopeCmd) (uuid.UUID, error) {
	// Check if scope already exists for this app
	if exists, _ := s.clientScopeRepo.Check.ByAppIDAndScope(ctx, cmd.AppID, cmd.Scope); exists {
		return uuid.Nil, clientScopeDomain.ErrClientScopeAlreadyExists
	}

	// Parse expires_at if provided
	var expiresAt *time.Time
	if cmd.ExpiresAt != nil && *cmd.ExpiresAt != "" {
		parsed, err := time.Parse(time.RFC3339, *cmd.ExpiresAt)
		if err != nil {
			return uuid.Nil, err
		}
		expiresAt = &parsed
	}

	// Create domain entity
	clientScope, err := clientScopeDomain.NewClientScope(clientScopeDomain.NewClientScopeParams{
		AppID:     cmd.AppID,
		Scope:     cmd.Scope,
		GrantedBy: cmd.GrantedBy,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.clientScopeRepo.Create.New(ctx, clientScope); err != nil {
		return uuid.Nil, err
	}

	return clientScope.ID(), nil
}
