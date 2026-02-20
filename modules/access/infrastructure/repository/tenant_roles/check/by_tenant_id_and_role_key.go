package check

import (
	"context"
	"nfxid/modules/access/infrastructure/rdb/models"
	"github.com/google/uuid"
)

func (h *Handler) ByTenantIDAndRoleKey(ctx context.Context, tenantID uuid.UUID, roleKey string) (bool, error) {
	var n int64
	err := h.db.WithContext(ctx).Model(&models.TenantRole{}).
		Where("tenant_id = ? AND role_key = ?", tenantID, roleKey).Count(&n).Error
	return n > 0, err
}
