package queries

import (
	imageDomain "nfxid/modules/image/domain/image"
	"nfxid/pkgs/query"
	"nfxid/pkgs/utils/ptr"
)

// ToDomainListQuery 将 ImageListQuery 转换为 domain.ListQuery
func (q ImageListQuery) ToDomainListQuery() imageDomain.ListQuery {
	domainQuery := imageDomain.ListQuery{
		TypeID:       q.TypeID,
		UserID:       q.UserID,
		SourceDomain: q.SourceDomain,
		IsPublic:     q.IsPublic,
		Search:       ptr.PtrIfNotZero(q.Search),
	}

	// Convert pagination
	if q.Page > 0 && q.PageSize > 0 {
		domainQuery.Offset = (q.Page - 1) * q.PageSize
		domainQuery.Limit = q.PageSize
	}

	// Convert sorting
	if q.OrderBy != "" {
		var sortField imageDomain.SortField
		switch q.OrderBy {
		case "created_at":
			sortField = imageDomain.SortByCreatedTime
		case "updated_at":
			sortField = imageDomain.SortByUpdatedAt
		case "size":
			sortField = imageDomain.SortBySize
		default:
			sortField = imageDomain.SortByCreatedTime
		}
		order := q.Order
		if order == "" {
			order = "desc"
		}
		domainQuery.DomainSorts = []query.DomainSort[imageDomain.SortField]{
			{Field: sortField, Order: order},
		}
	}

	return domainQuery
}
