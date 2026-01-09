package update

import (
	"context"
	"time"
	"nfxid/modules/tenants/domain/tenant_apps"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/tenant_apps/mapper"

	"github.com/google/uuid"
)

// Status 更新状态，实现 tenant_apps.Update 接口
func (h *Handler) Status(ctx context.Context, id uuid.UUID, status tenant_apps.TenantAppStatus) error {
	statusEnum := mapper.TenantAppStatusDomainToEnum(status)
	updates := map[string]any{
		models.TenantAppCols.Status:    statusEnum,
		models.TenantAppCols.UpdatedAt: time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.TenantApp{}).
		Where("id = ?", id).
		Updates(updates).Error
}
