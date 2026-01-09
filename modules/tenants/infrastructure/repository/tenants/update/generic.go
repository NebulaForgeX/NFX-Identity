package update

import (
	"context"
	"nfxid/modules/tenants/domain/tenants"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/tenants/mapper"
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
