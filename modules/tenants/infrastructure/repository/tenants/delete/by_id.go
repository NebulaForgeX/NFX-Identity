package delete

import (
	"context"
	"nfxid/modules/tenants/domain/tenants"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 Tenant，实现 tenants.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Tenant{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return tenants.ErrTenantNotFound
	}
	return nil
}
