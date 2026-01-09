package get

import (
	"context"
	"errors"
	"nfxid/modules/image/domain/images"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/images/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PublicByUserID 根据 UserID 获取公开的 Images，实现 images.Get 接口
func (h *Handler) PublicByUserID(ctx context.Context, userID uuid.UUID) ([]*images.Image, error) {
	var ms []models.Image
	if err := h.db.WithContext(ctx).Where("user_id = ? AND is_public = ?", userID, true).Find(&ms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*images.Image{}, nil
		}
		return nil, err
	}

	result := make([]*images.Image, len(ms))
	for i := range ms {
		result[i] = mapper.ImageModelToDomain(&ms[i])
	}
	return result, nil
}
