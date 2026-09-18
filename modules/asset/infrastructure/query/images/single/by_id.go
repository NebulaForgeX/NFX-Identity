package single

import (
	"context"
	"errors"

	asseterrs "nfxidentity/errors/src/asset"
	"nfxidentity/modules/asset/infrastructure/query/images/mapper"
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	imgs "nfxidentity/modules/asset/query/images"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*imgs.ImageVO, error) {
	var row rdbviews.ImagesActiveView
	if err := h.db.WithContext(ctx).
		Table(rdbviews.ImagesActiveView{}.TableName()).
		Where(rdbviews.ImagesActiveViewCols.ID+" = ?", id.String()).
		First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, asseterrs.ErrImageNotFound
		}
		return nil, asseterrs.ErrImageQueryFailed
	}
	vo := mapper.ImagesActiveViewToVO(&row)
	return &vo, nil
}
