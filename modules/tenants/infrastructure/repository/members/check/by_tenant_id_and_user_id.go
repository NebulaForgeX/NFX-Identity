package check

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByTenantIDAndUserID 根据 TenantID 和 UserID 检查 Member 是否存在，实现 members.Check 接口
func (h *Handler) ByTenantIDAndUserID(ctx context.Context, tenantID, userID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.Member{}).
		Where("tenant_id = ? AND user_id = ?", tenantID, userID).
		Count(&count).Error
	return count > 0, err
}
