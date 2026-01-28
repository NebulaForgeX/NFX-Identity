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

// ByID 根据 ID 获取 UserImage，实现 user_images.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*user_images.UserImage, error) {
	var m models.UserImage
	if err := h.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_images.ErrUserImageNotFound
		}
		return nil, err
	}
	return mapper.UserImageModelToDomain(&m), nil
}
