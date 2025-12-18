package list

import (
	"context"
	profileDomain "nfxid/modules/auth/domain/profile"
	profileDomainViews "nfxid/modules/auth/domain/profile/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/views"
)

// Generic 获取 Profile 列表，实现 profileDomain.List 接口
func (h *Handler) Generic(ctx context.Context, listQuery profileDomain.ListQuery) ([]*profileDomainViews.ProfileView, int64, error) {
	var items []views.ProfileCompleteView
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.ProfileListQueryToCommonQuery(listQuery)

	queryBuilder := h.db.WithContext(ctx).Model(&views.ProfileCompleteView{})

	// Apply search
	if commonQuery.Search != "" {
		queryBuilder = queryBuilder.Where(
			"display_name ILIKE ? OR nickname ILIKE ? OR bio ILIKE ?",
			"%"+commonQuery.Search+"%",
			"%"+commonQuery.Search+"%",
			"%"+commonQuery.Search+"%",
		)
	}

	// Count total
	if err := queryBuilder.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	if !commonQuery.All {
		if commonQuery.Offset > 0 {
			queryBuilder = queryBuilder.Offset(commonQuery.Offset)
		}
		if commonQuery.Limit > 0 {
			queryBuilder = queryBuilder.Limit(commonQuery.Limit)
		}
	}

	// Apply sorting
	if len(commonQuery.Sorts) > 0 {
		for _, sort := range commonQuery.Sorts {
			queryBuilder = queryBuilder.Order(sort.Field + " " + sort.Order)
		}
	} else {
		queryBuilder = queryBuilder.Order("profile_created_at DESC")
	}

	// Execute query
	if err := queryBuilder.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	// Convert to domain views
	result := make([]*profileDomainViews.ProfileView, len(items))
	for i, item := range items {
		view := mapper.ProfileViewToDomain(&item)
		result[i] = &view
	}

	return result, total, nil
}
