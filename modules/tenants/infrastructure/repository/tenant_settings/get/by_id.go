package get

import (
	"context"
	"errors"
	tenantsErr "nfxidentity/errors/src/tenants"
	"nfxidentity/modules/tenants/domain/tenant_settings"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/tenant_settings/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 TenantSetting，实现 tenant_settings.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*tenant_settings.TenantSetting, error) {
	var m models.TenantSetting
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenantsErr.ErrTenantSettingNotFound
		}
		return nil, err
	}
	return mapper.TenantSettingModelToDomain(&m), nil
}
