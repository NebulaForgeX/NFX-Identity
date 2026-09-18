package list

import (
	"context"

	asseterrs "nfxidentity/errors/src/asset"
	"nfxidentity/modules/asset/infrastructure/query/audios/mapper"
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	auds "nfxidentity/modules/asset/query/audios"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/query"
	"nfxidentity/pkgs/slicex"

	"gorm.io/gorm"
)

func (h *Handler) Generic(ctx context.Context, q auds.ListQuery) (httpx.Page[auds.AudioVO], error) {
	q.Normalize()
	params := mapper.AudioListQueryToParams(q)
	base := h.db.WithContext(ctx).Table(rdbviews.AudiosActiveView{}.TableName())

	rows, total, err := query.ExecuteQuery(
		ctx,
		base,
		params,
		videoQueryConfig,
		func(db *gorm.DB, data *[]rdbviews.AudiosActiveView) error {
			if len(params.Sorts) == 0 {
				db = db.Order(rdbviews.AudiosActiveViewCols.CreatedAt + " DESC")
			}
			return db.Find(data).Error
		},
	)
	if err != nil {
		return httpx.Page[auds.AudioVO]{}, asseterrs.ErrAudioQueryFailed
	}

	return httpx.NewPage(slicex.MapP(rows, mapper.AudiosActiveViewToVO), total), nil
}
