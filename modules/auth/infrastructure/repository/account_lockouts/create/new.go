package create

import (
	"context"
	"nfxid/modules/auth/domain/account_lockouts"
	"nfxid/modules/auth/infrastructure/repository/account_lockouts/mapper"
)

// New 创建新的 AccountLockout，实现 account_lockouts.Create 接口
func (h *Handler) New(ctx context.Context, al *account_lockouts.AccountLockout) error {
	m := mapper.AccountLockoutDomainToModel(al)
	return h.db.WithContext(ctx).Create(&m).Error
}
