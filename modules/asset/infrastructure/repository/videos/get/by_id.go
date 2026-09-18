package get

import (
	"context"
	"errors"

	asseterrs "nfxidentity/errors/src/asset"
	videosDomain "nfxidentity/modules/asset/domain/videos"
	rdbmodels "nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/modules/asset/infrastructure/repository/videos/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*videosDomain.Video, error) {
	var m rdbmodels.Video
	if err := h.db.WithContext(ctx).Where(rdbmodels.VideoCols.ID+" = ?", id.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, asseterrs.ErrVideoNotFound
		}
		return nil, err
	}
	return mapper.VideoModelToDomain(&m), nil
}
