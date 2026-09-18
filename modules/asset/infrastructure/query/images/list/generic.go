package list

import (
	"context"

	asseterrs "nfxidentity/errors/src/asset"
	"nfxidentity/modules/asset/infrastructure/query/images/mapper"
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	imgs "nfxidentity/modules/asset/query/images"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/query"
	"nfxidentity/pkgs/slicex"

	"gorm.io/gorm"
)

func (h *Handler) Generic(ctx context.Context, q imgs.ListQuery) (httpx.Page[imgs.ImageVO], error) {
	q.Normalize()
	params := mapper.ListQueryToParams(q)
	base := h.db.WithContext(ctx).Table(rdbviews.ImagesActiveView{}.TableName())

	rows, total, err := query.ExecuteQuery(
		ctx,
		base,
		params,
		imageQueryConfig,
		func(db *gorm.DB, data *[]rdbviews.ImagesActiveView) error {
			if len(params.Sorts) == 0 {
				db = db.Order(rdbviews.ImagesActiveViewCols.CreatedAt + " DESC")
			}
			return db.Find(data).Error
		},
	)
	if err != nil {
		return httpx.Page[imgs.ImageVO]{}, asseterrs.ErrImageQueryFailed
	}

	return httpx.NewPage(slicex.MapP(rows, mapper.ImagesActiveViewToVO), total), nil
}
