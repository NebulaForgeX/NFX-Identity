package get

import (
	"context"
	"errors"
	imageDomain "nfxid/modules/image/domain/image"
	imageDomainErrors "nfxid/modules/image/domain/image/errors"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/mapper"

	"gorm.io/gorm"
)

// ByFilename 根据 Filename 获取 Image，实现 imageDomain.Get 接口
func (h *Handler) ByFilename(ctx context.Context, filename string) (*imageDomain.Image, error) {
	var m models.Image
	if err := h.db.WithContext(ctx).Where("filename = ?", filename).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageDomainErrors.ErrImageNotFound
		}
		return nil, err
	}
	return mapper.ImageModelToDomain(&m), nil
}
