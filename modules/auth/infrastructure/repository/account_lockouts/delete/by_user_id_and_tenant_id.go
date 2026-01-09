package delete

import (
	"context"
	"nfxid/modules/auth/domain/account_lockouts"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserIDAndTenantID 根据 UserID 和 TenantID 删除 AccountLockout，实现 account_lockouts.Delete 接口
func (h *Handler) ByUserIDAndTenantID(ctx context.Context, userID, tenantID uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("user_id = ? AND tenant_id = ?", userID, tenantID).
		Delete(&models.AccountLockout{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return account_lockouts.ErrAccountLockoutNotFound
	}
	return nil
}
