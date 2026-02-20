package check

import (
	"context"

	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) ByUserIDAndTenantID(ctx context.Context, userID, tenantID uuid.UUID) (bool, error) {
	var n int64
	err := h.db.WithContext(ctx).Model(&models.TenantRoleAssignment{}).
		Where("user_id = ? AND tenant_id = ?", userID, tenantID).Count(&n).Error
	return n > 0, err
}
