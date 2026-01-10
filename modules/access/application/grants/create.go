package grants

import (
	"context"
	"time"
	grantCommands "nfxid/modules/access/application/grants/commands"
	grantDomain "nfxid/modules/access/domain/grants"

	"github.com/google/uuid"
)

// CreateGrant 创建授权
func (s *Service) CreateGrant(ctx context.Context, cmd grantCommands.CreateGrantCmd) (uuid.UUID, error) {
	var expiresAt *time.Time
	if cmd.ExpiresAt != nil && *cmd.ExpiresAt != "" {
		parsed, err := time.Parse(time.RFC3339, *cmd.ExpiresAt)
		if err != nil {
			return uuid.Nil, err
		}
		expiresAt = &parsed
	}

	// Create domain entity
	grant, err := grantDomain.NewGrant(grantDomain.NewGrantParams{
		SubjectType:  cmd.SubjectType,
		SubjectID:    cmd.SubjectID,
		GrantType:    cmd.GrantType,
		GrantRefID:   cmd.GrantRefID,
		TenantID:     cmd.TenantID,
		AppID:        cmd.AppID,
		ResourceType: cmd.ResourceType,
		ResourceID:   cmd.ResourceID,
		Effect:       cmd.Effect,
		ExpiresAt:    expiresAt,
		CreatedBy:    cmd.CreatedBy,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.grantRepo.Create.New(ctx, grant); err != nil {
		return uuid.Nil, err
	}

	return grant.ID(), nil
}
