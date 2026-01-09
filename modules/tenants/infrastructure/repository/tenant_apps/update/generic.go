package update

import (
	"context"
	"nfxid/modules/tenants/domain/tenant_apps"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/tenant_apps/mapper"
)

// Generic 通用更新 TenantApp，实现 tenant_apps.Update 接口
func (h *Handler) Generic(ctx context.Context, ta *tenant_apps.TenantApp) error {
	m := mapper.TenantAppDomainToModel(ta)
	updates := mapper.TenantAppModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.TenantApp{}).
		Where("id = ?", ta.ID()).
		Updates(updates).Error
}
