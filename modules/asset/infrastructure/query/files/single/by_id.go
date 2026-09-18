package single

import (
	"context"
	"errors"

	asseterrs "nfxidentity/errors/src/asset"
	"nfxidentity/modules/asset/infrastructure/query/files/mapper"
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	fls "nfxidentity/modules/asset/query/files"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*fls.FileVO, error) {
	var row rdbviews.FilesActiveView
	if err := h.db.WithContext(ctx).
		Table(rdbviews.FilesActiveView{}.TableName()).
		Where(rdbviews.FilesActiveViewCols.ID+" = ?", id.String()).
		First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, asseterrs.ErrFileNotFound
		}
		return nil, asseterrs.ErrFileQueryFailed
	}
	vo := mapper.FilesActiveViewToVO(&row)
	return &vo, nil
}
