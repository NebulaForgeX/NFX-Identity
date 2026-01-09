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

// ByUserID 根据 UserID 获取 AccountLockout，实现 account_lockouts.Get 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) (*account_lockouts.AccountLockout, error) {
	var m models.AccountLockout
	if err := h.db.WithContext(ctx).Where("user_id = ?", userID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, account_lockouts.ErrAccountLockoutNotFound
		}
		return nil, err
	}
	return mapper.AccountLockoutModelToDomain(&m), nil
}
