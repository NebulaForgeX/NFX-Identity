package update

import (
	"context"
	"errors"
	"time"
	"nfxid/modules/auth/domain/account_lockouts"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Unlock 解锁账户，实现 account_lockouts.Update 接口
func (h *Handler) Unlock(ctx context.Context, userID uuid.UUID, unlockedBy string, unlockActorID *uuid.UUID) error {
	// 先检查 AccountLockout 是否存在
	var m models.AccountLockout
	if err := h.db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return account_lockouts.ErrAccountLockoutNotFound
		}
		return err
	}

	// 检查是否已经解锁
	if m.UnlockedAt != nil {
		return account_lockouts.ErrAccountNotLocked
	}

	now := time.Now().UTC()
	updates := map[string]any{
		models.AccountLockoutCols.UnlockedAt:    &now,
		models.AccountLockoutCols.UnlockedBy:    &unlockedBy,
		models.AccountLockoutCols.UnlockActorID: unlockActorID,
		models.AccountLockoutCols.UpdatedAt:     now,
	}

	return h.db.WithContext(ctx).
		Model(&models.AccountLockout{}).
		Where("user_id = ?", userID).
		Updates(updates).Error
}
