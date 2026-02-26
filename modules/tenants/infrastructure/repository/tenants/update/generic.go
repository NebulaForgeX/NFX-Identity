package update

import (
	"context"
	"nfxidentity/modules/tenants/domain/tenants"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/tenants/mapper"
)

// Generic 通用更新 Tenant，实现 tenants.Update 接口
func (h *Handler) Generic(ctx context.Context, t *tenants.Tenant) error {
	m := mapper.TenantDomainToModel(t)
	updates := mapper.TenantModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Tenant{}).
		Where("id = ?", t.ID()).
		Updates(updates).Error
}
