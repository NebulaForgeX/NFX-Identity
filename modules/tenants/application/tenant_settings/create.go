package tenant_settings

import (
	"context"
	tenantSettingCommands "nfxid/modules/tenants/application/tenant_settings/commands"
	tenantSettingDomain "nfxid/modules/tenants/domain/tenant_settings"

	"github.com/google/uuid"
)

// CreateTenantSetting 创建租户设置
func (s *Service) CreateTenantSetting(ctx context.Context, cmd tenantSettingCommands.CreateTenantSettingCmd) (uuid.UUID, error) {
	// Check if tenant setting already exists
	if exists, _ := s.tenantSettingRepo.Check.ByTenantID(ctx, cmd.TenantID); exists {
		return uuid.Nil, tenantSettingDomain.ErrTenantSettingAlreadyExists
	}

	// Create domain entity
	tenantSetting, err := tenantSettingDomain.NewTenantSetting(tenantSettingDomain.NewTenantSettingParams{
		TenantID:           cmd.TenantID,
		EnforceMFA:         cmd.EnforceMFA,
		AllowedEmailDomains: cmd.AllowedEmailDomains,
		SessionTTLMinutes:  cmd.SessionTTLMinutes,
		PasswordPolicy:     cmd.PasswordPolicy,
		LoginPolicy:        cmd.LoginPolicy,
		MFAPolicy:          cmd.MFAPolicy,
	})
	if err != nil {
		return uuid.Nil, err
	}

	// Save to repository
	if err := s.tenantSettingRepo.Create.New(ctx, tenantSetting); err != nil {
		return uuid.Nil, err
	}

	return tenantSetting.ID(), nil
}
