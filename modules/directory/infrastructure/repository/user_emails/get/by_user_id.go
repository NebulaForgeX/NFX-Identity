package get

import (
	"context"
	"nfxid/modules/directory/domain/user_emails"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_emails/mapper"

	"github.com/google/uuid"
)

// ByUserID 根据 UserID 获取 UserEmail 列表，实现 user_emails.Get 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) ([]*user_emails.UserEmail, error) {
	var ms []models.UserEmail
	if err := h.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*user_emails.UserEmail, len(ms))
	for i, m := range ms {
		result[i] = mapper.UserEmailModelToDomain(&m)
	}
	return result, nil
}
