package tenants

import (
	"context"
	tenantCommands "nfxid/modules/tenants/application/tenants/commands"
)

// DeleteTenant 删除租户（软删除）
func (s *Service) DeleteTenant(ctx context.Context, cmd tenantCommands.DeleteTenantCmd) error {
	// Get domain entity
	tenant, err := s.tenantRepo.Get.ByID(ctx, cmd.TenantID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := tenant.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.tenantRepo.Update.Generic(ctx, tenant)
}
