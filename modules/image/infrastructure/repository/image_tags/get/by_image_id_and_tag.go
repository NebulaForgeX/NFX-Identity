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

// ByImageIDAndTag 根据 ImageID 和 Tag 获取 ImageTag，实现 image_tags.Get 接口
func (h *Handler) ByImageIDAndTag(ctx context.Context, imageID uuid.UUID, tag string) (*image_tags.ImageTag, error) {
	var m models.ImageTag
	if err := h.db.WithContext(ctx).
		Where("image_id = ? AND tag = ?", imageID, tag).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, image_tags.ErrImageTagNotFound
		}
		return nil, err
	}
	return mapper.ImageTagModelToDomain(&m), nil
}
