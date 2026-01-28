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

// CurrentByUserID 根据 UserID 获取当前 UserImage（display_order = 0），实现 user_images.Get 接口
func (h *Handler) CurrentByUserID(ctx context.Context, userID uuid.UUID) (*user_images.UserImage, error) {
	var m models.UserImage
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND display_order = 0 AND deleted_at IS NULL", userID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_images.ErrUserImageNotFound
		}
		return nil, err
	}
	return mapper.UserImageModelToDomain(&m), nil
}
