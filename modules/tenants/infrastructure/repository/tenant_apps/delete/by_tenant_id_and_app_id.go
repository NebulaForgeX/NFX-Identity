package delete

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByTenantIDAndAppID 根据 TenantID 和 AppID 删除 TenantApp，实现 tenant_apps.Delete 接口
func (h *Handler) ByTenantIDAndAppID(ctx context.Context, tenantID, appID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where("tenant_id = ? AND application_id = ?", tenantID, appID).
		Delete(&models.TenantApplication{}).Error
}
