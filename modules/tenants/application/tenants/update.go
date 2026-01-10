package tenants

import (
	"context"
	tenantCommands "nfxid/modules/tenants/application/tenants/commands"
)

// UpdateTenant 更新租户
func (s *Service) UpdateTenant(ctx context.Context, cmd tenantCommands.UpdateTenantCmd) error {
	// Get domain entity
	tenant, err := s.tenantRepo.Get.ByID(ctx, cmd.TenantID)
	if err != nil {
		return err
	}

	// Update domain entity
	if err := tenant.Update(cmd.Name, cmd.DisplayName, cmd.PrimaryDomain, cmd.Metadata); err != nil {
		return err
	}

	// Save to repository
	return s.tenantRepo.Update.Generic(ctx, tenant)
}

// UpdateTenantStatus 更新租户状态
func (s *Service) UpdateTenantStatus(ctx context.Context, cmd tenantCommands.UpdateTenantStatusCmd) error {
	// Get domain entity
	tenant, err := s.tenantRepo.Get.ByID(ctx, cmd.TenantID)
	if err != nil {
		return err
	}

	// Update status domain entity
	if err := tenant.UpdateStatus(cmd.Status); err != nil {
		return err
	}

	// Save to repository
	return s.tenantRepo.Update.Status(ctx, cmd.TenantID, cmd.Status)
}
