package get

import (
	"context"
	"errors"

	imageErr "nfxidentity/errors/src/image"
	"nfxidentity/modules/image/domain/images"
	"nfxidentity/modules/image/infrastructure/rdb/models"
	"nfxidentity/modules/image/infrastructure/repository/images/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Image，实现 images.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*images.Image, error) {
	var m models.Image
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageErr.ErrImageNotFound
		}
		return nil, err
	}
	return mapper.ImageModelToDomain(&m), nil
}
