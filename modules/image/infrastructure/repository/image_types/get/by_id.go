package get

import (
	"context"
	"errors"

	imageErr "nfxidentity/errors/src/image"
	"nfxidentity/modules/image/domain/image_types"
	"nfxidentity/modules/image/infrastructure/rdb/models"
	"nfxidentity/modules/image/infrastructure/repository/image_types/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 ImageType，实现 image_types.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*image_types.ImageType, error) {
	var m models.ImageType
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageErr.ErrImageTypeNotFound
		}
		return nil, err
	}
	return mapper.ImageTypeModelToDomain(&m), nil
}
