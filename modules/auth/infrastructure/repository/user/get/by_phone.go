package get

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/user"
	userDomainErrors "nfxid/modules/auth/domain/user/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"gorm.io/gorm"
)

// ByPhone 根据 Phone 获取 User，实现 user.Get 接口
func (h *Handler) ByPhone(ctx context.Context, phone string) (*user.User, error) {
	var m models.User
	if err := h.db.WithContext(ctx).Where("phone = ?", phone).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userDomainErrors.ErrUserNotFound
		}
		return nil, err
	}
	return mapper.UserModelToDomain(&m), nil
}
