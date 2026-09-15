package get

import (
	"context"
	"errors"

	"nfxidentity/modules/asset/domain/images"
	"nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/modules/asset/infrastructure/repository/images/mapper"
	"nfxidentity/pkgs/errx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*images.Image, error) {
	var m models.Image
	if err := h.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NotFound("ASSET_NOT_FOUND", "asset not found")
		}
		return nil, err
	}
	return mapper.ToDomain(&m), nil
}
