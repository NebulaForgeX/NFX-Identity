package tenants

import (
	"context"
	tenantCommands "nfxid/modules/tenants/application/tenants/commands"
	tenantDomain "nfxid/modules/tenants/domain/tenants"

	"github.com/google/uuid"
)

// CreateTenant 创建租户
func (s *Service) CreateTenant(ctx context.Context, cmd tenantCommands.CreateTenantCmd) (uuid.UUID, error) {
	// Check if tenant ID already exists
	if exists, _ := s.tenantRepo.Check.ByTenantID(ctx, cmd.TenantID); exists {
		return uuid.Nil, tenantDomain.ErrTenantIDAlreadyExists
	}

	// Create domain entity
	tenant, err := tenantDomain.NewTenant(tenantDomain.NewTenantParams{
		TenantID:      cmd.TenantID,
		Name:          cmd.Name,
		DisplayName:   cmd.DisplayName,
		Status:        cmd.Status,
		PrimaryDomain: cmd.PrimaryDomain,
		Metadata:      cmd.Metadata,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.tenantRepo.Create.New(ctx, tenant); err != nil {
		return uuid.Nil, err
	}

	return tenant.ID(), nil
}
