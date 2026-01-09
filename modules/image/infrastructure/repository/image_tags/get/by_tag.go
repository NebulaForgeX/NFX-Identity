package get

import (
	"context"
	"errors"
	"nfxid/modules/image/domain/image_tags"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/image_tags/mapper"

	"gorm.io/gorm"
)

// ByTag 根据 Tag 获取 ImageTags，实现 image_tags.Get 接口
func (h *Handler) ByTag(ctx context.Context, tag string) ([]*image_tags.ImageTag, error) {
	var ms []models.ImageTag
	if err := h.db.WithContext(ctx).Where("tag = ?", tag).Find(&ms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*image_tags.ImageTag{}, nil
		}
		return nil, err
	}
	
	result := make([]*image_tags.ImageTag, len(ms))
	for i := range ms {
		result[i] = mapper.ImageTagModelToDomain(&ms[i])
	}
	return result, nil
}
