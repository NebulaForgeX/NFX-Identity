package update

import (
	"context"
	"time"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// UpdateVerificationCode 更新验证码，实现 user_phones.Update 接口
func (h *Handler) UpdateVerificationCode(ctx context.Context, id uuid.UUID, code string, expiresAt time.Time) error {
	updates := map[string]any{
		models.UserPhoneCols.VerificationCode:      &code,
		models.UserPhoneCols.VerificationExpiresAt: &expiresAt,
		models.UserPhoneCols.UpdatedAt:             time.Now().UTC(),
	}

	return h.db.WithContext(ctx).
		Model(&models.UserPhone{}).
		Where("id = ?", id).
		Updates(updates).Error
}
