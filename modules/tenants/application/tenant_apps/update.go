package tenant_apps

import (
	"context"
	tenantAppCommands "nfxid/modules/tenants/application/tenant_apps/commands"
)

// UpdateTenantAppStatus 更新租户应用状态
func (s *Service) UpdateTenantAppStatus(ctx context.Context, cmd tenantAppCommands.UpdateTenantAppStatusCmd) error {
	// Get domain entity
	tenantApp, err := s.tenantAppRepo.Get.ByID(ctx, cmd.TenantAppID)
	if err != nil {
		return err
	}

	// Update status domain entity
	if err := tenantApp.UpdateStatus(cmd.Status); err != nil {
		return err
	}

	// Save to repository
	return s.tenantAppRepo.Update.Status(ctx, cmd.TenantAppID, cmd.Status)
}

// UpdateTenantAppSettings 更新租户应用设置
func (s *Service) UpdateTenantAppSettings(ctx context.Context, cmd tenantAppCommands.UpdateTenantAppSettingsCmd) error {
	// Get domain entity
	tenantApp, err := s.tenantAppRepo.Get.ByID(ctx, cmd.TenantAppID)
	if err != nil {
		return err
	}

	// Update settings domain entity
	if err := tenantApp.UpdateSettings(cmd.Settings); err != nil {
		return err
	}

	// Save to repository
	return s.tenantAppRepo.Update.Generic(ctx, tenantApp)
}
