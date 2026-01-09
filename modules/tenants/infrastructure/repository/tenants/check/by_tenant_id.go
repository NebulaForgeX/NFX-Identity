package check

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"
)

// ByTenantID 根据 TenantID 检查 Tenant 是否存在，实现 tenants.Check 接口
func (h *Handler) ByTenantID(ctx context.Context, tenantID string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.Tenant{}).
		Where("tenant_id = ?", tenantID).
		Count(&count).Error
	return count > 0, err
}
