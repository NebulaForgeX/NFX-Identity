package get

import (
	"context"
	"errors"
	imageTypeDomain "nfxid/modules/image/domain/image_type"
	imageTypeDomainErrors "nfxid/modules/image/domain/image_type/errors"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 ImageType，实现 imageTypeDomain.Get 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*imageTypeDomain.ImageType, error) {
	var m models.ImageType
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, imageTypeDomainErrors.ErrImageTypeNotFound
		}
		return nil, err
	}
	return mapper.ImageTypeModelToDomain(&m), nil
}
