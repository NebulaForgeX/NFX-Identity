package update

import (
	"context"
	"nfxid/modules/tenants/domain/tenant_settings"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/tenant_settings/mapper"
)

// Generic 通用更新 TenantSetting，实现 tenant_settings.Update 接口
func (h *Handler) Generic(ctx context.Context, ts *tenant_settings.TenantSetting) error {
	m := mapper.TenantSettingDomainToModel(ts)
	updates := mapper.TenantSettingModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.TenantSetting{}).
		Where("id = ?", ts.ID()).
		Updates(updates).Error
}
