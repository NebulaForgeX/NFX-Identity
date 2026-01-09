package delete

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"
)

// ByEmail 根据 Email 删除 UserEmail，实现 user_emails.Delete 接口
func (h *Handler) ByEmail(ctx context.Context, email string) error {
	return h.db.WithContext(ctx).
		Where("email = ?", email).
		Delete(&models.UserEmail{}).Error
}
