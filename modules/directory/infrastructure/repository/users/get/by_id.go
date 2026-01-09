package get

import (
	"context"
	"errors"
	"nfxid/modules/directory/domain/users"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/users/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 User，实现 users.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*users.User, error) {
	var m models.User
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, users.ErrUserNotFound
		}
		return nil, err
	}
	return mapper.UserModelToDomain(&m), nil
}
