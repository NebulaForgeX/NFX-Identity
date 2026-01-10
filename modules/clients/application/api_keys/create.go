package api_keys

import (
	"context"
	"time"
	apiKeyCommands "nfxid/modules/clients/application/api_keys/commands"
	apiKeyDomain "nfxid/modules/clients/domain/api_keys"

	"github.com/google/uuid"
)

// CreateAPIKey 创建API密钥
func (s *Service) CreateAPIKey(ctx context.Context, cmd apiKeyCommands.CreateAPIKeyCmd) (uuid.UUID, error) {
	// Check if key_id already exists
	if exists, _ := s.apiKeyRepo.Check.ByKeyID(ctx, cmd.KeyID); exists {
		return uuid.Nil, apiKeyDomain.ErrKeyIDAlreadyExists
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
	apiKey, err := apiKeyDomain.NewAPIKey(apiKeyDomain.NewAPIKeyParams{
		KeyID:     cmd.KeyID,
		AppID:     cmd.AppID,
		KeyHash:   cmd.KeyHash,
		HashAlg:   cmd.HashAlg,
		Name:      cmd.Name,
		Status:    apiKeyDomain.APIKeyStatusActive,
		ExpiresAt: expiresAt,
		CreatedBy: cmd.CreatedBy,
		Metadata:  cmd.Metadata,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.apiKeyRepo.Create.New(ctx, apiKey); err != nil {
		return uuid.Nil, err
	}

	return apiKey.ID(), nil
}
