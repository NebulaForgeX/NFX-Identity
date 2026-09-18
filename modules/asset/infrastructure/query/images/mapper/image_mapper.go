package mapper

import (
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	imgs "nfxidentity/modules/asset/query/images"
	"nfxidentity/pkgs/ptrx"
	"nfxidentity/pkgs/query"
	"nfxidentity/pkgs/slicex"
)

func ImagesActiveViewToVO(v *rdbviews.ImagesActiveView) imgs.ImageVO {
	if v == nil {
		return imgs.ImageVO{}
	}
	return imgs.ImageVO{
		ID:         ptrx.Deref(v.ID),
		FilePath:   ptrx.Deref(v.FilePath),
		FileName:   ptrx.Deref(v.FileName),
		FileSize:   ptrx.Deref(v.FileSize),
		MimeType:   ptrx.Deref(v.MimeType),
		Width:      v.Width,
		Height:     v.Height,
		AltText:    v.AltText,
		UploaderID: ptrx.Deref(v.UploaderID),
		CreatedAt:  ptrx.Deref(v.CreatedAt),
		UpdatedAt:  ptrx.Deref(v.UpdatedAt),
	}
}

var sortFieldToCol = map[imgs.SortField]string{
	imgs.SortByCreatedAt: rdbviews.ImagesActiveViewCols.CreatedAt,
	imgs.SortByUpdatedAt: rdbviews.ImagesActiveViewCols.UpdatedAt,
}

func ListQueryToParams(q imgs.ListQuery) *query.ListQueryParams {
	search := ""
	if q.Search != nil {
		search = *q.Search
	}
	cols := rdbviews.ImagesActiveViewCols
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
		Sorts:   query.DomainSortToSort(q.DomainSorts, sortFieldToCol),
		Filters: filters,
	}
}
