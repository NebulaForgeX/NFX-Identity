package delete

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 TenantApp，实现 tenant_apps.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.TenantApp{}).Error
}
