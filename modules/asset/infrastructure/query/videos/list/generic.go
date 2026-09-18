package list

import (
	"context"

	asseterrs "nfxidentity/errors/src/asset"
	"nfxidentity/modules/asset/infrastructure/query/videos/mapper"
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	vids "nfxidentity/modules/asset/query/videos"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/query"
	"nfxidentity/pkgs/slicex"

	"gorm.io/gorm"
)

func (h *Handler) Generic(ctx context.Context, q vids.ListQuery) (httpx.Page[vids.VideoVO], error) {
	q.Normalize()
	params := mapper.VideoListQueryToParams(q)
	base := h.db.WithContext(ctx).Table(rdbviews.VideosActiveView{}.TableName())

	rows, total, err := query.ExecuteQuery(
		ctx,
		base,
		params,
		videoQueryConfig,
		func(db *gorm.DB, data *[]rdbviews.VideosActiveView) error {
			if len(params.Sorts) == 0 {
				db = db.Order(rdbviews.VideosActiveViewCols.CreatedAt + " DESC")
			}
			return db.Find(data).Error
		},
	)
	if err != nil {
		return httpx.Page[vids.VideoVO]{}, asseterrs.ErrVideoQueryFailed
	}

	return httpx.NewPage(slicex.MapP(rows, mapper.VideosActiveViewToVO), total), nil
}
