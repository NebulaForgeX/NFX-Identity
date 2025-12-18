package list

import (
	"context"
	permissionDomain "nfxid/modules/permission/domain/permission"
	permissionDomainViews "nfxid/modules/permission/domain/permission/views"
	"nfxid/modules/permission/infrastructure/rdb/models"
	"nfxid/pkgs/query"
	"nfxid/pkgs/utils/ptr"
	"nfxid/pkgs/utils/slice"

	"gorm.io/gorm"
)

var permissionQueryConfig = &query.QueryConfig{
	SearchConfig: &query.SearchConfig{
		Fields:   []string{"tag", "name", "description"},
		Operator: "ILIKE",
		Logic:    "OR",
	},
}

// Generic 获取 Permission 列表，实现 permissionDomain.List 接口
func (h *Handler) Generic(ctx context.Context, listQuery permissionDomain.ListQuery) ([]*permissionDomainViews.PermissionView, int64, error) {
	var items []models.Permission
	var total int64

	itemsResult, total, err := query.ExecuteQuery(
		ctx,
		h.db.WithContext(ctx).Model(&models.Permission{}).Where("deleted_at IS NULL"),
		permissionListQueryToCommonQuery(listQuery),
		permissionQueryConfig,
		func(db *gorm.DB, data *[]models.Permission) error {
			// Apply additional filters
			if listQuery.Category != nil {
				db = db.Where("category = ?", *listQuery.Category)
			}
			if listQuery.IsSystem != nil {
				db = db.Where("is_system = ?", *listQuery.IsSystem)
			}
			if len(listQuery.PermissionIDs) > 0 {
				db = db.Where("id IN ?", listQuery.PermissionIDs)
			}
			return db.Find(data).Error
		},
	)
	if err != nil {
		return nil, 0, err
	}
	items = itemsResult

	// Convert to pointers
	result := make([]*permissionDomainViews.PermissionView, len(items))
	for i, item := range items {
		view := permissionModelToDomainView(&item)
		result[i] = &view
	}
	return result, total, nil
}

func permissionListQueryToCommonQuery(q permissionDomain.ListQuery) *query.ListQueryParams {
	filters := make(map[string][]any)
	if len(q.PermissionIDs) > 0 {
		filters["id"] = slice.ToAnySlice(q.PermissionIDs)
	}
	if q.Category != nil {
		filters["category"] = []any{*q.Category}
	}
	if q.IsSystem != nil {
		filters["is_system"] = []any{*q.IsSystem}
	}

	return &query.ListQueryParams{
		All:    q.All,
		Offset: q.Offset,
		Limit:  q.Limit,
		Sorts: query.DomainSortToSort(q.DomainSorts, map[permissionDomain.SortField]string{
			permissionDomain.SortByCreatedTime: "created_at",
			permissionDomain.SortByTag:         "tag",
			permissionDomain.SortByCategory:    "category",
		}),
		Search:  ptr.Deref(q.Search),
		Filters: filters,
	}
}
