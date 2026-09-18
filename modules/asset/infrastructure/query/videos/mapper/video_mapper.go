package mapper

import (
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	vids "nfxidentity/modules/asset/query/videos"
	"nfxidentity/pkgs/ptrx"
	"nfxidentity/pkgs/query"
	"nfxidentity/pkgs/slicex"
)

func VideosActiveViewToVO(v *rdbviews.VideosActiveView) vids.VideoVO {
	if v == nil {
		return vids.VideoVO{}
	}
	return vids.VideoVO{
		ID:              ptrx.Deref(v.ID),
		FilePath:        ptrx.Deref(v.FilePath),
		FileName:        ptrx.Deref(v.FileName),
		FileSize:        ptrx.Deref(v.FileSize),
		MimeType:        ptrx.Deref(v.MimeType),
		DurationSeconds: v.DurationSeconds,
		Width:           v.Width,
		Height:          v.Height,
		UploaderID:      ptrx.Deref(v.UploaderID),
		CreatedAt:       ptrx.Deref(v.CreatedAt),
		UpdatedAt:       ptrx.Deref(v.UpdatedAt),
	}
}

var videoSortFieldToCol = map[vids.SortField]string{
	vids.SortByCreatedAt: rdbviews.VideosActiveViewCols.CreatedAt,
	vids.SortByUpdatedAt: rdbviews.VideosActiveViewCols.UpdatedAt,
}

func VideoListQueryToParams(q vids.ListQuery) *query.ListQueryParams {
	search := ""
	if q.Search != nil {
		search = *q.Search
	}
	cols := rdbviews.VideosActiveViewCols
	filters := map[string][]any{}
	if len(q.IDs) > 0 {
		filters[cols.ID] = slicex.ToAnySlice(slicex.UuidSliceToStrSlice(q.IDs))
	}
	if len(q.UploaderIDs) > 0 {
		filters[cols.UploaderID] = slicex.ToAnySlice(slicex.UuidSliceToStrSlice(q.UploaderIDs))
	}
	if len(q.MimeTypes) > 0 {
		filters[cols.MimeType] = slicex.ToAnySlice(q.MimeTypes)
	}
	return &query.ListQueryParams{
		Offset:  q.DomainPagination.Offset,
		Limit:   q.DomainPagination.Limit,
		All:     q.DomainPagination.All,
		Search:  search,
		Sorts:   query.DomainSortToSort(q.DomainSorts, videoSortFieldToCol),
		Filters: filters,
	}
}
