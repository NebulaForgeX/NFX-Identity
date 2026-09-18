package mapper

import (
	rdbviews "nfxidentity/modules/asset/infrastructure/rdb/views"
	fls "nfxidentity/modules/asset/query/files"
	"nfxidentity/pkgs/ptrx"
	"nfxidentity/pkgs/query"
	"nfxidentity/pkgs/slicex"
)

func FilesActiveViewToVO(v *rdbviews.FilesActiveView) fls.FileVO {
	if v == nil {
		return fls.FileVO{}
	}
	return fls.FileVO{
		ID:         ptrx.Deref(v.ID),
		FilePath:   ptrx.Deref(v.FilePath),
		FileName:   ptrx.Deref(v.FileName),
		FileSize:   ptrx.Deref(v.FileSize),
		MimeType:   ptrx.Deref(v.MimeType),
		UploaderID: ptrx.Deref(v.UploaderID),
		CreatedAt:  ptrx.Deref(v.CreatedAt),
		UpdatedAt:  ptrx.Deref(v.UpdatedAt),
	}
}

var fileSortFieldToCol = map[fls.SortField]string{
	fls.SortByCreatedAt: rdbviews.FilesActiveViewCols.CreatedAt,
	fls.SortByUpdatedAt: rdbviews.FilesActiveViewCols.UpdatedAt,
}

func FileListQueryToParams(q fls.ListQuery) *query.ListQueryParams {
	search := ""
	if q.Search != nil {
		search = *q.Search
	}
	cols := rdbviews.FilesActiveViewCols
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
		Sorts:   query.DomainSortToSort(q.DomainSorts, fileSortFieldToCol),
		Filters: filters,
	}
}
