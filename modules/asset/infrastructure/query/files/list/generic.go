package list

import (
	"context"

	asseterrs "nfxidentity/errors/src/asset"
	"nfxidentity/modules/asset/infrastructure/query/files/mapper"
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	fls "nfxidentity/modules/asset/query/files"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/query"
	"nfxidentity/pkgs/slicex"

	"gorm.io/gorm"
)

func (h *Handler) Generic(ctx context.Context, q fls.ListQuery) (httpx.Page[fls.FileVO], error) {
	q.Normalize()
	params := mapper.FileListQueryToParams(q)
	base := h.db.WithContext(ctx).Table(rdbviews.FilesActiveView{}.TableName())

	rows, total, err := query.ExecuteQuery(
		ctx,
		base,
		params,
		fileQueryConfig,
		func(db *gorm.DB, data *[]rdbviews.FilesActiveView) error {
			if len(params.Sorts) == 0 {
				db = db.Order(rdbviews.FilesActiveViewCols.CreatedAt + " DESC")
			}
			return db.Find(data).Error
		},
	)
	if err != nil {
		return httpx.Page[fls.FileVO]{}, asseterrs.ErrFileQueryFailed
	}

	return httpx.NewPage(slicex.MapP(rows, mapper.FilesActiveViewToVO), total), nil
}
