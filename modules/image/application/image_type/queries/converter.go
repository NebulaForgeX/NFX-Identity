package queries

import (
	imageTypeDomain "nfxid/modules/image/domain/image_type"
	"nfxid/pkgs/query"
	"nfxid/pkgs/utils/ptr"
)

// ToDomainListQuery 将 ImageTypeListQuery 转换为 domain.ListQuery
func (q ImageTypeListQuery) ToDomainListQuery() imageTypeDomain.ListQuery {
	domainQuery := imageTypeDomain.ListQuery{
		IsSystem: q.IsSystem,
		Search:   ptr.PtrIfNotZero(q.Search),
	}

	// Convert pagination
	if q.Page > 0 && q.PageSize > 0 {
		domainQuery.Offset = (q.Page - 1) * q.PageSize
		domainQuery.Limit = q.PageSize
	}

	// Convert sorting
	if q.OrderBy != "" {
		var sortField imageTypeDomain.SortField
		switch q.OrderBy {
		case "created_at":
			sortField = imageTypeDomain.SortByCreatedTime
		case "key":
			sortField = imageTypeDomain.SortByKey
		default:
			sortField = imageTypeDomain.SortByCreatedTime
		}
		order := q.Order
		if order == "" {
			order = "desc"
		}
		domainQuery.DomainSorts = []query.DomainSort[imageTypeDomain.SortField]{
			{Field: sortField, Order: order},
		}
	}

	return domainQuery
}
