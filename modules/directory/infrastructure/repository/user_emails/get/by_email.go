package get

import (
	"context"
	"errors"
	"nfxid/modules/directory/domain/user_emails"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_emails/mapper"

	"gorm.io/gorm"
)

// ByEmail 根据 Email 获取 UserEmail，实现 user_emails.Get 接口
func (h *Handler) ByEmail(ctx context.Context, email string) (*user_emails.UserEmail, error) {
	var m models.UserEmail
	if err := h.db.WithContext(ctx).Where("email = ?", email).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_emails.ErrUserEmailNotFound
		}
		return nil, err
	}
	return mapper.UserEmailModelToDomain(&m), nil
}
