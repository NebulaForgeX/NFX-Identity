package get

import (
	"context"
	"errors"
	"nfxid/modules/directory/domain/user_phones"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_phones/mapper"

	"gorm.io/gorm"
)

// ByPhone 根据 Phone 获取 UserPhone，实现 user_phones.Get 接口
func (h *Handler) ByPhone(ctx context.Context, phone string) (*user_phones.UserPhone, error) {
	var m models.UserPhone
	if err := h.db.WithContext(ctx).Where("phone = ?", phone).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_phones.ErrUserPhoneNotFound
		}
		return nil, err
	}
	return mapper.UserPhoneModelToDomain(&m), nil
}
