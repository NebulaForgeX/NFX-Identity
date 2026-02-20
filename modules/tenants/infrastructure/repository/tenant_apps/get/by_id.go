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

// ByID 根据 ID 获取 TenantApp，实现 tenant_apps.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*tenant_apps.TenantApp, error) {
	var m models.TenantApplication
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenant_apps.ErrTenantAppNotFound
		}
		return nil, err
	}
	return mapper.TenantAppModelToDomain(&m), nil
}
