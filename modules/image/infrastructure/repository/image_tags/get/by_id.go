package get

import (
	"context"
	"errors"
	imageErr "nfxidentity/errors/src/image"
	"nfxidentity/modules/image/domain/image_tags"
	"nfxidentity/modules/image/infrastructure/rdb/models"
	"nfxidentity/modules/image/infrastructure/repository/image_tags/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 ImageTag，实现 image_tags.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*image_tags.ImageTag, error) {
	var m models.ImageTag
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageErr.ErrImageTagNotFound
		}
		return nil, err
	}
	return mapper.ImageTagModelToDomain(&m), nil
}
