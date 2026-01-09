package get

import (
	"context"
	"errors"
	"nfxid/modules/image/domain/image_tags"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/image_tags/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 ImageTag，实现 image_tags.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*image_tags.ImageTag, error) {
	var m models.ImageTag
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, image_tags.ErrImageTagNotFound
		}
		return nil, err
	}
	return mapper.ImageTagModelToDomain(&m), nil
}
