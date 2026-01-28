package get

import (
	"context"
	"errors"
	"nfxid/modules/directory/domain/user_images"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_images/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByUserIDAndImageID 根据 UserID 和 ImageID 获取 UserImage，实现 user_images.Get 接口
func (h *Handler) ByUserIDAndImageID(ctx context.Context, userID, imageID uuid.UUID) (*user_images.UserImage, error) {
	var m models.UserImage
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND image_id = ? AND deleted_at IS NULL", userID, imageID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_images.ErrUserImageNotFound
		}
		return nil, err
	}
	return mapper.UserImageModelToDomain(&m), nil
}
