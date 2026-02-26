package create

import (
	"context"
	"nfxidentity/modules/tenants/domain/tenants"
	"nfxidentity/modules/tenants/infrastructure/repository/tenants/mapper"
)

// New 创建新的 Tenant，实现 tenants.Create 接口
func (h *Handler) New(ctx context.Context, t *tenants.Tenant) error {
	m := mapper.TenantDomainToModel(t)
	return h.db.WithContext(ctx).Create(&m).Error
}
