package tenant_apps

import (
	"context"
	tenantAppCommands "nfxid/modules/tenants/application/tenant_apps/commands"
)

// DeleteTenantApp 删除租户应用
func (s *Service) DeleteTenantApp(ctx context.Context, cmd tenantAppCommands.DeleteTenantAppCmd) error {
	// Delete from repository (hard delete)
	return s.tenantAppRepo.Delete.ByID(ctx, cmd.TenantAppID)
}
