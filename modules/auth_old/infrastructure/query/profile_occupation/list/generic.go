package list

import (
	"context"
	occupationDomain "nfxid/modules/auth/domain/profile_occupation"
	occupationDomainViews "nfxid/modules/auth/domain/profile_occupation/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/slice"
)

// Generic 获取 Occupation 列表，实现 occupationDomain.List 接口
func (h *Handler) Generic(ctx context.Context, listQuery occupationDomain.ListQuery) ([]*occupationDomainViews.OccupationView, int64, error) {
	var items []models.Occupation
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.OccupationListQueryToCommonQuery(listQuery)

	queryBuilder := h.db.WithContext(ctx).Model(&models.Occupation{}).Where("deleted_at IS NULL")

	// Apply search
	if commonQuery.Search != "" {
		queryBuilder = queryBuilder.Where(
			"company ILIKE ? OR position ILIKE ? OR department ILIKE ? OR industry ILIKE ?",
			"%"+commonQuery.Search+"%",
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
		queryBuilder = queryBuilder.Order("start_date DESC")
	}

	// Execute query
	if err := queryBuilder.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	result := slice.MapP(items, mapper.OccupationModelToDomain)
	// Convert to pointers
	pointerResult := make([]*occupationDomainViews.OccupationView, len(result))
	for i := range result {
		pointerResult[i] = &result[i]
	}
	return pointerResult, total, nil
}
