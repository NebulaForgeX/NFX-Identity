package update

import (
	"context"
	"time"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// Verify 验证用户，实现 users.Update 接口
func (h *Handler) Verify(ctx context.Context, id uuid.UUID) error {
	updates := map[string]any{
		models.UserCols.IsVerified: true,
		models.UserCols.UpdatedAt:   time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(updates).Error
}
