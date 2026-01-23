package check

import (
	"context"
	"time"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// IsLocked 检查账户是否被锁定，实现 account_lockouts.Check 接口
func (h *Handler) IsLocked(ctx context.Context, userID uuid.UUID) (bool, error) {
	var m models.AccountLockout
	if err := h.db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&m).Error; err != nil {
		return false, err
	}

	// 检查是否有锁定时间，且锁定时间未过期
	if m.LockedUntil != nil {
		return m.LockedUntil.After(time.Now()), nil
	}

	// 如果没有锁定时间，检查是否有解锁时间
	return m.UnlockedAt == nil, nil
}
