package update

import (
	"context"
	"nfxid/modules/directory/domain/user_emails"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_emails/mapper"
)

// Generic 通用更新 UserEmail，实现 user_emails.Update 接口
func (h *Handler) Generic(ctx context.Context, ue *user_emails.UserEmail) error {
	m := mapper.UserEmailDomainToModel(ue)
	updates := mapper.UserEmailModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.UserEmail{}).
		Where("id = ?", ue.ID()).
		Updates(updates).Error
}
