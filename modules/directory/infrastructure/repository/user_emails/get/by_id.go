package get

import (
	"context"
	"errors"
	"nfxid/modules/directory/domain/user_emails"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_emails/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 UserEmail，实现 user_emails.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*user_emails.UserEmail, error) {
	var m models.UserEmail
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_emails.ErrUserEmailNotFound
		}
		return nil, err
	}
	return mapper.UserEmailModelToDomain(&m), nil
}
