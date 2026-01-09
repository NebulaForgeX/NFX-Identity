package get

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/tenant_apps"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/tenant_apps/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByTenantIDAndAppID 根据 TenantID 和 AppID 获取 TenantApp，实现 tenant_apps.Get 接口
func (h *Handler) ByTenantIDAndAppID(ctx context.Context, tenantID, appID uuid.UUID) (*tenant_apps.TenantApp, error) {
	var m models.TenantApp
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ? AND app_id = ?", tenantID, appID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenant_apps.ErrTenantAppNotFound
		}
		return nil, err
	}
	return mapper.TenantAppModelToDomain(&m), nil
}
