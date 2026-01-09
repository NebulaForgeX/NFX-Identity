package update

import (
	"context"
	"time"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// SetPrimary 设置为主邮箱，实现 user_emails.Update 接口
func (h *Handler) SetPrimary(ctx context.Context, id uuid.UUID) error {
	// 先获取这个邮箱的 user_id
	var m models.UserEmail
	if err := h.db.WithContext(ctx).
		Where("id = ?", id).
		First(&m).Error; err != nil {
		return err
	}

	// 先将该用户的所有邮箱设置为非主邮箱
	updates1 := map[string]any{
		models.UserEmailCols.IsPrimary: false,
		models.UserEmailCols.UpdatedAt: time.Now().UTC(),
	}
	if err := h.db.WithContext(ctx).
		Model(&models.UserEmail{}).
		Where("user_id = ?", m.UserID).
		Updates(updates1).Error; err != nil {
		return err
	}

	// 然后设置这个邮箱为主邮箱
	updates2 := map[string]any{
		models.UserEmailCols.IsPrimary: true,
		models.UserEmailCols.UpdatedAt: time.Now().UTC(),
	}
	return h.db.WithContext(ctx).
		Model(&models.UserEmail{}).
		Where("id = ?", id).
		Updates(updates2).Error
}
