package list

import (
	"context"
	"encoding/json"
	imageDomain "nfxid/modules/image/domain/image"
	imageDomainViews "nfxid/modules/image/domain/image/views"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/pkgs/query"
	"nfxid/pkgs/utils/ptr"
	"nfxid/pkgs/utils/slice"

	"gorm.io/gorm"
)

var imageQueryConfig = &query.QueryConfig{
	SearchConfig: &query.SearchConfig{
		Fields:   []string{"filename", "original_filename"},
		Operator: "ILIKE",
		Logic:    "OR",
	},
}

// Generic 获取 Image 列表，实现 imageDomain.List 接口
func (h *Handler) Generic(ctx context.Context, listQuery imageDomain.ListQuery) ([]*imageDomainViews.ImageView, int64, error) {
	var items []models.Image
	var total int64

	itemsResult, total, err := query.ExecuteQuery(
		ctx,
		h.db.WithContext(ctx).Model(&models.Image{}),
		imageListQueryToCommonQuery(listQuery),
		imageQueryConfig,
		func(db *gorm.DB, data *[]models.Image) error {
			// Apply additional filters
			if listQuery.TypeID != nil {
				db = db.Where("type_id = ?", *listQuery.TypeID)
			}
			if listQuery.UserID != nil {
				db = db.Where("user_id = ?", *listQuery.UserID)
			}
			if listQuery.SourceDomain != nil {
				db = db.Where("source_domain = ?", *listQuery.SourceDomain)
			}
			if listQuery.IsPublic != nil {
				db = db.Where("is_public = ?", *listQuery.IsPublic)
			}
			if len(listQuery.ImageIDs) > 0 {
				db = db.Where("id IN ?", listQuery.ImageIDs)
			}
			return db.Find(data).Error
		},
	)
	if err != nil {
		return nil, 0, err
	}
	items = itemsResult

	// Convert to pointers
	result := make([]*imageDomainViews.ImageView, len(items))
	for i, item := range items {
		view := imageModelToDomainView(&item)
		result[i] = &view
	}
	return result, total, nil
}

func imageListQueryToCommonQuery(q imageDomain.ListQuery) *query.ListQueryParams {
	filters := make(map[string][]any)
	if len(q.ImageIDs) > 0 {
		filters["id"] = slice.ToAnySlice(q.ImageIDs)
	}
	if q.TypeID != nil {
		filters["type_id"] = []any{*q.TypeID}
	}
	if q.UserID != nil {
		filters["user_id"] = []any{*q.UserID}
	}
	if q.SourceDomain != nil {
		filters["source_domain"] = []any{*q.SourceDomain}
	}
	if q.IsPublic != nil {
		filters["is_public"] = []any{*q.IsPublic}
	}

	return &query.ListQueryParams{
		All:    q.All,
		Offset: q.Offset,
		Limit:  q.Limit,
		Sorts: query.DomainSortToSort(q.DomainSorts, map[imageDomain.SortField]string{
			imageDomain.SortByCreatedTime: "created_at",
			imageDomain.SortByUpdatedAt:   "updated_at",
			imageDomain.SortBySize:        "size",
		}),
		Search:  ptr.Deref(q.Search),
		Filters: filters,
	}
}

func imageModelToDomainView(m *models.Image) imageDomainViews.ImageView {
	var metadata map[string]interface{}
	if m.Metadata != nil {
		rawJSON, _ := m.Metadata.MarshalJSON()
		if len(rawJSON) > 0 {
			_ = json.Unmarshal(rawJSON, &metadata)
		}
	}

	return imageDomainViews.ImageView{
		ID:               m.ID,
		TypeID:           m.TypeID,
		UserID:           m.UserID,
		SourceDomain:     m.SourceDomain,
		Filename:         m.Filename,
		OriginalFilename: m.OriginalFilename,
		MimeType:         m.MimeType,
		Size:             m.Size,
		Width:            m.Width,
		Height:           m.Height,
		StoragePath:      m.StoragePath,
		URL:              m.URL,
		IsPublic:         m.IsPublic,
		Metadata:         metadata,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
	}
}
