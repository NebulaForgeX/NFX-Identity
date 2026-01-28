package get

import (
	"context"
	"nfxid/modules/directory/domain/user_images"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_images/mapper"

	"github.com/google/uuid"
)

// ByImageID 根据 ImageID 获取 UserImage 列表，实现 user_images.Get 接口
func (h *Handler) ByImageID(ctx context.Context, imageID uuid.UUID) ([]*user_images.UserImage, error) {
	var ms []models.UserImage
	if err := h.db.WithContext(ctx).
		Where("image_id = ? AND deleted_at IS NULL", imageID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*user_images.UserImage, len(ms))
	for i, m := range ms {
		result[i] = mapper.UserImageModelToDomain(&m)
	}
	return result, nil
}
