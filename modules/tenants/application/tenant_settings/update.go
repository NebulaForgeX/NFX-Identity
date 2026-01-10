package tenant_settings

import (
	"context"
	tenantSettingCommands "nfxid/modules/tenants/application/tenant_settings/commands"
)

// UpdateTenantSetting 更新租户设置
func (s *Service) UpdateTenantSetting(ctx context.Context, cmd tenantSettingCommands.UpdateTenantSettingCmd) error {
	// Get domain entity
	tenantSetting, err := s.tenantSettingRepo.Get.ByID(ctx, cmd.TenantSettingID)
	if err != nil {
		return err
	}

	// Update domain entity
	if err := tenantSetting.Update(cmd.EnforceMFA, cmd.AllowedEmailDomains, cmd.SessionTTLMinutes, cmd.PasswordPolicy, cmd.LoginPolicy, cmd.MFAPolicy, cmd.UpdatedBy); err != nil {
		return err
	}

	// Save to repository
	return s.tenantSettingRepo.Update.Generic(ctx, tenantSetting)
}
