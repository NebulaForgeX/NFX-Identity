package update

import (
	"context"
	"nfxidentity/modules/tenants/domain/tenant_apps"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/tenant_apps/mapper"
	"time"

	"github.com/google/uuid"
)

// Status 更新状态，实现 tenant_apps.Update 接口
func (h *Handler) Status(ctx context.Context, id uuid.UUID, status tenant_apps.TenantAppStatus) error {
	statusEnum := mapper.TenantAppStatusDomainToEnum(status)
	updates := map[string]any{
		models.TenantApplicationCols.Status:    statusEnum,
		models.TenantApplicationCols.UpdatedAt: time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.TenantApplication{}).
		Where("id = ?", id).
		Updates(updates).Error
}
