package get

import (
	"context"
	"errors"
	"nfxid/modules/directory/domain/user_avatars"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_avatars/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByImageID 根据 ImageID 获取 UserAvatar，实现 user_avatars.Get 接口
func (h *Handler) ByImageID(ctx context.Context, imageID uuid.UUID) (*user_avatars.UserAvatar, error) {
	var m models.UserAvatar
	if err := h.db.WithContext(ctx).Where("image_id = ?", imageID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_avatars.ErrUserAvatarNotFound
		}
		return nil, err
	}
	return mapper.UserAvatarModelToDomain(&m), nil
}
