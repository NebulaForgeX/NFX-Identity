package get

import (
	"context"
	"errors"
	"nfxid/modules/image/domain/images"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/images/mapper"

	"gorm.io/gorm"
)

// BySourceDomain 根据 SourceDomain 获取 Images，实现 images.Get 接口
func (h *Handler) BySourceDomain(ctx context.Context, sourceDomain string) ([]*images.Image, error) {
	var ms []models.Image
	if err := h.db.WithContext(ctx).Where("source_domain = ?", sourceDomain).Find(&ms).Error; err != nil {
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
