package get

import (
	"context"
	"errors"
	tenantsErr "nfxidentity/errors/src/tenants"
	"nfxidentity/modules/tenants/domain/tenant_apps"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/tenant_apps/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByTenantIDAndAppID 根据 TenantID 和 AppID 获取 TenantApp，实现 tenant_apps.Get 接口
func (h *Handler) ByTenantIDAndAppID(ctx context.Context, tenantID, appID uuid.UUID) (*tenant_apps.TenantApp, error) {
	var m models.TenantApplication
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ? AND application_id = ?", tenantID, appID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenantsErr.ErrTenantAppNotFound
		}
		return nil, err
	}
	return mapper.TenantAppModelToDomain(&m), nil
}
