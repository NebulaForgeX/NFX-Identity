package count

import (
	"context"

	asseterrs "nfxidentity/errors/src/asset"
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
)

func (h *Handler) All(ctx context.Context) (int64, error) {
	var n int64
	if err := h.db.WithContext(ctx).Table(rdbviews.VideosActiveView{}.TableName()).Count(&n).Error; err != nil {
		return 0, asseterrs.ErrVideoQueryFailed
	}
	return n, nil
}
