package get

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/tenant_settings"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/tenant_settings/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByTenantID 根据 TenantID 获取 TenantSetting，实现 tenant_settings.Get 接口
func (h *Handler) ByTenantID(ctx context.Context, tenantID uuid.UUID) (*tenant_settings.TenantSetting, error) {
	var m models.TenantSetting
	if err := h.db.WithContext(ctx).Where("tenant_id = ?", tenantID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenant_settings.ErrTenantSettingNotFound
		}
		return nil, err
	}
	return mapper.TenantSettingModelToDomain(&m), nil
}
