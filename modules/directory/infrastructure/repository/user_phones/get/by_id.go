package get

import (
	"context"
	"errors"
	"nfxid/modules/directory/domain/user_phones"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_phones/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 UserPhone，实现 user_phones.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*user_phones.UserPhone, error) {
	var m models.UserPhone
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_phones.ErrUserPhoneNotFound
		}
		return nil, err
	}
	return mapper.UserPhoneModelToDomain(&m), nil
}
