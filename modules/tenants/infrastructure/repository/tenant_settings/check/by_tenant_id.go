package check

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByTenantID 根据 TenantID 检查 TenantSetting 是否存在，实现 tenant_settings.Check 接口
func (h *Handler) ByTenantID(ctx context.Context, tenantID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.TenantSetting{}).
		Where("id = ?", tenantID).
		Count(&count).Error
	return count > 0, err
}
