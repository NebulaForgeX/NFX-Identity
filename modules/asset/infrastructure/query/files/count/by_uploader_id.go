package count

import (
	"context"

	asseterrs "nfxidentity/errors/src/asset"
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"

	"github.com/google/uuid"
)

func (h *Handler) ByUploaderID(ctx context.Context, uploaderID uuid.UUID) (int64, error) {
	var n int64
	if err := h.db.WithContext(ctx).
		Table(rdbviews.FilesActiveView{}.TableName()).
		Where(rdbviews.FilesActiveViewCols.UploaderID+" = ?", uploaderID.String()).
		Count(&n).Error; err != nil {
		return 0, asseterrs.ErrFileQueryFailed
	}
	return n, nil
}
