package create

import (
	"context"
	"nfxid/modules/directory/domain/user_emails"
	"nfxid/modules/directory/infrastructure/repository/user_emails/mapper"
)

// New 创建新的 UserEmail，实现 user_emails.Create 接口
func (h *Handler) New(ctx context.Context, ue *user_emails.UserEmail) error {
	m := mapper.UserEmailDomainToModel(ue)
	return h.db.WithContext(ctx).Create(&m).Error
}
