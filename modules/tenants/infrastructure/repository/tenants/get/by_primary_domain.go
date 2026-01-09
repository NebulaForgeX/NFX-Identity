package get

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/tenants"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/tenants/mapper"

	"gorm.io/gorm"
)

// ByPrimaryDomain 根据 PrimaryDomain 获取 Tenant，实现 tenants.Get 接口
func (h *Handler) ByPrimaryDomain(ctx context.Context, primaryDomain string) (*tenants.Tenant, error) {
	var m models.Tenant
	if err := h.db.WithContext(ctx).Where("primary_domain = ?", primaryDomain).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenants.ErrTenantNotFound
		}
		return nil, err
	}
	return mapper.TenantModelToDomain(&m), nil
}
