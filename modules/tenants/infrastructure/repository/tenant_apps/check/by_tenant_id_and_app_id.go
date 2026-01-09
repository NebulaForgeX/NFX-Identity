package check

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByTenantIDAndAppID 根据 TenantID 和 AppID 检查 TenantApp 是否存在，实现 tenant_apps.Check 接口
func (h *Handler) ByTenantIDAndAppID(ctx context.Context, tenantID, appID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.TenantApp{}).
		Where("tenant_id = ? AND app_id = ?", tenantID, appID).
		Count(&count).Error
	return count > 0, err
}
