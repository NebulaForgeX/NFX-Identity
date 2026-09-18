package single

import (
	"context"
	"errors"

	asseterrs "nfxidentity/errors/src/asset"
	"nfxidentity/modules/asset/infrastructure/query/videos/mapper"
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	vids "nfxidentity/modules/asset/query/videos"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*vids.VideoVO, error) {
	var row rdbviews.VideosActiveView
	if err := h.db.WithContext(ctx).
		Table(rdbviews.VideosActiveView{}.TableName()).
		Where(rdbviews.VideosActiveViewCols.ID+" = ?", id.String()).
		First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, asseterrs.ErrVideoNotFound
		}
		return nil, asseterrs.ErrVideoQueryFailed
	}
	vo := mapper.VideosActiveViewToVO(&row)
	return &vo, nil
}
