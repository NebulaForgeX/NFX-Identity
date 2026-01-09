package update

import (
	"context"
	"nfxid/modules/auth/domain/account_lockouts"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/account_lockouts/mapper"
)

// Generic 通用更新 AccountLockout，实现 account_lockouts.Update 接口
func (h *Handler) Generic(ctx context.Context, al *account_lockouts.AccountLockout) error {
	m := mapper.AccountLockoutDomainToModel(al)
	updates := mapper.AccountLockoutModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.AccountLockout{}).
		Where("user_id = ? AND tenant_id = ?", al.UserID(), al.TenantID()).
		Updates(updates).Error
}
