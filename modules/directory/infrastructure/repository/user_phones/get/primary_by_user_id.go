package get

import (
	"context"
	"errors"

	dirErr "nfxidentity/errors/src/directory"
	"nfxidentity/modules/directory/domain/user_phones"
	"nfxidentity/modules/directory/infrastructure/rdb/models"
	"nfxidentity/modules/directory/infrastructure/repository/user_phones/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PrimaryByUserID 根据 UserID 获取主手机号，实现 user_phones.Get 接口
func (h *Handler) PrimaryByUserID(ctx context.Context, userID uuid.UUID) (*user_phones.UserPhone, error) {
	var m models.UserPhone
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND is_primary = ?", userID, true).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dirErr.ErrUserPhoneNotFound
		}
		return nil, err
	}
	return mapper.UserPhoneModelToDomain(&m), nil
}
