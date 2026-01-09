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

// PrimaryByUserID 根据 UserID 获取主邮箱，实现 user_emails.Get 接口
func (h *Handler) PrimaryByUserID(ctx context.Context, userID uuid.UUID) (*user_emails.UserEmail, error) {
	var m models.UserEmail
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND is_primary = ?", userID, true).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_emails.ErrUserEmailNotFound
		}
		return nil, err
	}
	return mapper.UserEmailModelToDomain(&m), nil
}
