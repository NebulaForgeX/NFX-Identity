package delete

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/tenant_settings"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 删除 TenantSetting，实现 tenant_settings.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.TenantSetting{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return tenant_settings.ErrTenantSettingNotFound
	}
	return nil
}
