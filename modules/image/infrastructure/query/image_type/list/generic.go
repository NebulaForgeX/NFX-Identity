package list

import (
	"context"
	imageTypeDomain "nfxid/modules/image/domain/image_type"
	imageTypeDomainViews "nfxid/modules/image/domain/image_type/views"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/pkgs/query"
	"nfxid/pkgs/utils/ptr"
	"nfxid/pkgs/utils/slice"

	"gorm.io/gorm"
)

var imageTypeQueryConfig = &query.QueryConfig{
	SearchConfig: &query.SearchConfig{
		Fields:   []string{"key", "description"},
		Operator: "ILIKE",
		Logic:    "OR",
	},
}

// Generic 获取 ImageType 列表，实现 imageTypeDomain.List 接口
func (h *Handler) Generic(ctx context.Context, listQuery imageTypeDomain.ListQuery) ([]*imageTypeDomainViews.ImageTypeView, int64, error) {
	var items []models.ImageType
	var total int64

	itemsResult, total, err := query.ExecuteQuery(
		ctx,
		h.db.WithContext(ctx).Model(&models.ImageType{}),
		imageTypeListQueryToCommonQuery(listQuery),
		imageTypeQueryConfig,
		func(db *gorm.DB, data *[]models.ImageType) error {
			// Apply additional filters
			if listQuery.IsSystem != nil {
				db = db.Where("is_system = ?", *listQuery.IsSystem)
			}
			if len(listQuery.ImageTypeIDs) > 0 {
				db = db.Where("id IN ?", listQuery.ImageTypeIDs)
			}
			return db.Find(data).Error
		},
	)
	if err != nil {
		return nil, 0, err
	}
	items = itemsResult

	// Convert to pointers
	result := make([]*imageTypeDomainViews.ImageTypeView, len(items))
	for i, item := range items {
		view := imageTypeModelToDomainView(&item)
		result[i] = &view
	}
	return result, total, nil
}

func imageTypeListQueryToCommonQuery(q imageTypeDomain.ListQuery) *query.ListQueryParams {
	filters := make(map[string][]any)
	if len(q.ImageTypeIDs) > 0 {
		filters["id"] = slice.ToAnySlice(q.ImageTypeIDs)
	}
	if q.IsSystem != nil {
		filters["is_system"] = []any{*q.IsSystem}
	}

	return &query.ListQueryParams{
		All:    q.All,
		Offset: q.Offset,
		Limit:  q.Limit,
		Sorts: query.DomainSortToSort(q.DomainSorts, map[imageTypeDomain.SortField]string{
			imageTypeDomain.SortByCreatedTime: "created_at",
			imageTypeDomain.SortByKey:         "key",
		}),
		Search:  ptr.Deref(q.Search),
		Filters: filters,
	}
}

func imageTypeModelToDomainView(m *models.ImageType) imageTypeDomainViews.ImageTypeView {
	isSystem := false
	if m.IsSystem != nil {
		isSystem = *m.IsSystem
	}

	return imageTypeDomainViews.ImageTypeView{
		ID:          m.ID,
		Key:         m.Key,
		Description: m.Description,
		MaxWidth:    m.MaxWidth,
		MaxHeight:   m.MaxHeight,
		AspectRatio: m.AspectRatio,
		IsSystem:    isSystem,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
