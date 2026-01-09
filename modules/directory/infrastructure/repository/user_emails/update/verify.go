package update

import (
	"context"
	"time"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// Verify 验证邮箱，实现 user_emails.Update 接口
func (h *Handler) Verify(ctx context.Context, id uuid.UUID) error {
	now := time.Now().UTC()
	updates := map[string]any{
		models.UserEmailCols.IsVerified: true,
		models.UserEmailCols.VerifiedAt: &now,
		models.UserEmailCols.UpdatedAt:   now,
	}

	return h.db.WithContext(ctx).
		Model(&models.UserEmail{}).
		Where("id = ?", id).
		Updates(updates).Error
}
