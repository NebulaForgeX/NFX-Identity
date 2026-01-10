package tenant_apps

import (
	"context"
	tenantAppCommands "nfxid/modules/tenants/application/tenant_apps/commands"
	tenantAppDomain "nfxid/modules/tenants/domain/tenant_apps"

	"github.com/google/uuid"
)

// CreateTenantApp 创建租户应用
func (s *Service) CreateTenantApp(ctx context.Context, cmd tenantAppCommands.CreateTenantAppCmd) (uuid.UUID, error) {
	// Check if tenant app already exists
	if exists, _ := s.tenantAppRepo.Check.ByTenantIDAndAppID(ctx, cmd.TenantID, cmd.AppID); exists {
		return uuid.Nil, tenantAppDomain.ErrTenantAppAlreadyExists
	}

	// Create domain entity
	tenantApp, err := tenantAppDomain.NewTenantApp(tenantAppDomain.NewTenantAppParams{
		TenantID:  cmd.TenantID,
		AppID:     cmd.AppID,
		Status:    cmd.Status,
		CreatedBy: cmd.CreatedBy,
		Settings:  cmd.Settings,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.tenantAppRepo.Create.New(ctx, tenantApp); err != nil {
		return uuid.Nil, err
	}

	return tenantApp.ID(), nil
}
