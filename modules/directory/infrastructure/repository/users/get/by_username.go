package get

import (
	"context"
	"errors"
	"nfxid/modules/directory/domain/users"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/users/mapper"

	"gorm.io/gorm"
)

// ByUsername 根据 Username 获取 User，实现 users.Get 接口
func (h *Handler) ByUsername(ctx context.Context, username string) (*users.User, error) {
	var m models.User
	if err := h.db.WithContext(ctx).Where("username = ?", username).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, users.ErrUserNotFound
		}
		return nil, err
	}
	return mapper.UserModelToDomain(&m), nil
}
