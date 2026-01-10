package client_credentials

import (
	"context"
	"time"
	clientCredentialCommands "nfxid/modules/clients/application/client_credentials/commands"
	clientCredentialDomain "nfxid/modules/clients/domain/client_credentials"

	"github.com/google/uuid"
)

// CreateClientCredential 创建客户端凭证
func (s *Service) CreateClientCredential(ctx context.Context, cmd clientCredentialCommands.CreateClientCredentialCmd) (uuid.UUID, error) {
	// Check if client_id already exists
	if exists, _ := s.clientCredentialRepo.Check.ByClientID(ctx, cmd.ClientID); exists {
		return uuid.Nil, clientCredentialDomain.ErrClientIDAlreadyExists
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
	clientCredential, err := clientCredentialDomain.NewClientCredential(clientCredentialDomain.NewClientCredentialParams{
		AppID:      cmd.AppID,
		ClientID:   cmd.ClientID,
		SecretHash: cmd.SecretHash,
		HashAlg:    cmd.HashAlg,
		Status:     clientCredentialDomain.CredentialStatusActive,
		ExpiresAt:  expiresAt,
		CreatedBy:  cmd.CreatedBy,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.clientCredentialRepo.Create.New(ctx, clientCredential); err != nil {
		return uuid.Nil, err
	}

	return clientCredential.ID(), nil
}
