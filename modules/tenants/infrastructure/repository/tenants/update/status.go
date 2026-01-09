package update

import (
	"context"
	"time"
	"nfxid/modules/tenants/domain/tenants"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/tenants/mapper"

	"github.com/google/uuid"
)

// Status 更新状态，实现 tenants.Update 接口
func (h *Handler) Status(ctx context.Context, id uuid.UUID, status tenants.TenantStatus) error {
	statusEnum := mapper.TenantStatusDomainToEnum(status)
	updates := map[string]any{
		models.TenantCols.Status:    statusEnum,
		models.TenantCols.UpdatedAt: time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.Tenant{}).
		Where("id = ?", id).
		Updates(updates).Error
}
