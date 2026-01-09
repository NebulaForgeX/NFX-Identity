package get

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/account_lockouts"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/account_lockouts/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByUserIDAndTenantID 根据 UserID 和 TenantID 获取 AccountLockout，实现 account_lockouts.Get 接口
func (h *Handler) ByUserIDAndTenantID(ctx context.Context, userID, tenantID uuid.UUID) (*account_lockouts.AccountLockout, error) {
	var m models.AccountLockout
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND tenant_id = ?", userID, tenantID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, account_lockouts.ErrAccountLockoutNotFound
		}
		return nil, err
	}
	return mapper.AccountLockoutModelToDomain(&m), nil
}
