package get

import (
	"context"
	"errors"

	asseterrs "nfxidentity/errors/src/asset"
	audiosDomain "nfxidentity/modules/asset/domain/audios"
	rdbmodels "nfxidentity/modules/asset/infrastructure/rdb/models"
	"nfxidentity/modules/asset/infrastructure/repository/audios/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*audiosDomain.Audio, error) {
	var m rdbmodels.Audio
	if err := h.db.WithContext(ctx).Where(rdbmodels.AudioCols.ID+" = ?", id.String()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, asseterrs.ErrAudioNotFound
		}
		return nil, err
	}
	return mapper.AudioModelToDomain(&m), nil
}
