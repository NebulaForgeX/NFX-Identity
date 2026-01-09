package check

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserIDAndTenantID 根据 UserID 和 TenantID 检查 AccountLockout 是否存在，实现 account_lockouts.Check 接口
func (h *Handler) ByUserIDAndTenantID(ctx context.Context, userID, tenantID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.AccountLockout{}).
		Where("user_id = ? AND tenant_id = ?", userID, tenantID).
		Count(&count).Error
	return count > 0, err
}
