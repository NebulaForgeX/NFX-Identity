package create

import (
	"context"
	"nfxid/modules/tenants/domain/tenant_settings"
	"nfxid/modules/tenants/infrastructure/repository/tenant_settings/mapper"
)

// New 创建新的 TenantSetting，实现 tenant_settings.Create 接口
func (h *Handler) New(ctx context.Context, ts *tenant_settings.TenantSetting) error {
	m := mapper.TenantSettingDomainToModel(ts)
	return h.db.WithContext(ctx).Create(&m).Error
}
