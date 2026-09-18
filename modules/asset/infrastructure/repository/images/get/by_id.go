package get

import (
	"context"
	"errors"
	"nfxidentity/errors/src/asset"

	"nfxidentity/modules/asset/domain/images"
	"nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/modules/asset/infrastructure/repository/images/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*images.Image, error) {
	var m models.Image
	if err := h.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, asset.ErrAssetNotFound
		}
		return nil, err
	}
	return mapper.ToDomain(&m), nil
}
