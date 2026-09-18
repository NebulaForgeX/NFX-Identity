package mapper

import (
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	auds "nfxidentity/modules/asset/query/audios"
	"nfxidentity/pkgs/ptrx"
	"nfxidentity/pkgs/query"
	"nfxidentity/pkgs/slicex"
)

func AudiosActiveViewToVO(v *rdbviews.AudiosActiveView) auds.AudioVO {
	if v == nil {
		return auds.AudioVO{}
	}
	return auds.AudioVO{
		ID:              ptrx.Deref(v.ID),
		FilePath:        ptrx.Deref(v.FilePath),
		FileName:        ptrx.Deref(v.FileName),
		FileSize:        ptrx.Deref(v.FileSize),
		MimeType:        ptrx.Deref(v.MimeType),
		DurationSeconds: v.DurationSeconds,
		UploaderID:      ptrx.Deref(v.UploaderID),
		CreatedAt:       ptrx.Deref(v.CreatedAt),
		UpdatedAt:       ptrx.Deref(v.UpdatedAt),
	}
}

var audioSortFieldToCol = map[auds.SortField]string{
	auds.SortByCreatedAt: rdbviews.AudiosActiveViewCols.CreatedAt,
	auds.SortByUpdatedAt: rdbviews.AudiosActiveViewCols.UpdatedAt,
}

func AudioListQueryToParams(q auds.ListQuery) *query.ListQueryParams {
	search := ""
	if q.Search != nil {
		search = *q.Search
	}
	cols := rdbviews.AudiosActiveViewCols
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
		Sorts:   query.DomainSortToSort(q.DomainSorts, audioSortFieldToCol),
		Filters: filters,
	}
}
