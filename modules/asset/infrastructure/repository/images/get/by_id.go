package get

import (
	"context"
	"errors"

	asseterrs "nfxidentity/errors/src/asset"
	imagesDomain "nfxidentity/modules/asset/domain/images"
	rdbmodels "nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/modules/asset/infrastructure/repository/images/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*imagesDomain.Image, error) {
	var m rdbmodels.Image
	if err := h.db.WithContext(ctx).Where(rdbmodels.ImageCols.ID+" = ?", id.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, asseterrs.ErrImageNotFound
		}
		return nil, err
	}
	return mapper.ImageModelToDomain(&m), nil
}
