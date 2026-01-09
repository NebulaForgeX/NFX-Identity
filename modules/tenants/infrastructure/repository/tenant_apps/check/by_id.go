package check

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 检查 TenantApp 是否存在，实现 tenant_apps.Check 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.TenantApp{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}
