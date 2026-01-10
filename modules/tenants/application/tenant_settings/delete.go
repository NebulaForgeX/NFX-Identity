package tenant_settings

import (
	"context"
	tenantSettingCommands "nfxid/modules/tenants/application/tenant_settings/commands"
)

// DeleteTenantSetting 删除租户设置
func (s *Service) DeleteTenantSetting(ctx context.Context, cmd tenantSettingCommands.DeleteTenantSettingCmd) error {
	// Delete from repository (hard delete)
	return s.tenantSettingRepo.Delete.ByID(ctx, cmd.TenantSettingID)
}
