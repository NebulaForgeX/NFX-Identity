package update

import (
	"context"
	"time"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// Verify 验证手机号，实现 user_phones.Update 接口
func (h *Handler) Verify(ctx context.Context, id uuid.UUID) error {
	now := time.Now().UTC()
	updates := map[string]any{
		models.UserPhoneCols.IsVerified: true,
		models.UserPhoneCols.VerifiedAt: &now,
		models.UserPhoneCols.UpdatedAt:   now,
	}

	return h.db.WithContext(ctx).
		Model(&models.UserPhone{}).
		Where("id = ?", id).
		Updates(updates).Error
}
