package single

import (
	"context"
	"errors"

	asseterrs "nfxidentity/errors/src/asset"
	"nfxidentity/modules/asset/infrastructure/query/audios/mapper"
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	auds "nfxidentity/modules/asset/query/audios"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*auds.AudioVO, error) {
	var row rdbviews.AudiosActiveView
	if err := h.db.WithContext(ctx).
		Table(rdbviews.AudiosActiveView{}.TableName()).
		Where(rdbviews.AudiosActiveViewCols.ID+" = ?", id.String()).
		First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, asseterrs.ErrAudioNotFound
		}
		return nil, asseterrs.ErrAudioQueryFailed
	}
	vo := mapper.AudiosActiveViewToVO(&row)
	return &vo, nil
}
