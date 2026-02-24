package get

import (
	"context"
	"errors"

	dirErr "nfxid/errors/src/directory"
	"nfxid/modules/directory/domain/user_avatars"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_avatars/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByUserID 根据 UserID 获取 UserAvatar，实现 user_avatars.Get 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) (*user_avatars.UserAvatar, error) {
	var m models.UserAvatar
	if err := h.db.WithContext(ctx).Where("user_id = ?", userID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dirErr.ErrUserAvatarNotFound
		}
		return nil, err
	}
	return mapper.UserAvatarModelToDomain(&m), nil
}
