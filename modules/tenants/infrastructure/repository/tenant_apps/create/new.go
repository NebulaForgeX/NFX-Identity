package create

import (
	"context"
	"nfxid/modules/tenants/domain/tenant_apps"
	"nfxid/modules/tenants/infrastructure/repository/tenant_apps/mapper"
)

// New 创建新的 TenantApp，实现 tenant_apps.Create 接口
func (h *Handler) New(ctx context.Context, ta *tenant_apps.TenantApp) error {
	m := mapper.TenantAppDomainToModel(ta)
	return h.db.WithContext(ctx).Create(&m).Error
}
