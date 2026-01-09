package get

import (
	"context"
	"nfxid/modules/directory/domain/user_phones"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_phones/mapper"

	"github.com/google/uuid"
)

// ByUserID 根据 UserID 获取 UserPhone 列表，实现 user_phones.Get 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) ([]*user_phones.UserPhone, error) {
	var ms []models.UserPhone
	if err := h.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*user_phones.UserPhone, len(ms))
	for i, m := range ms {
		result[i] = mapper.UserPhoneModelToDomain(&m)
	}
	return result, nil
}
