package update

import (
	"context"
	"time"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// UpdateLastLogin 更新最后登录时间，实现 users.Update 接口
func (h *Handler) UpdateLastLogin(ctx context.Context, id uuid.UUID) error {
	now := time.Now().UTC()
	updates := map[string]any{
		models.UserCols.LastLoginAt: &now,
		models.UserCols.UpdatedAt:    now,
	}

	return h.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(updates).Error
}
