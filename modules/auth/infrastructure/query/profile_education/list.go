package profile_education

import (
	"context"
	educationDomain "nfxid/modules/auth/domain/profile_education"
	educationDomainViews "nfxid/modules/auth/domain/profile_education/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/slice"
)

// List 获取 Education 列表，实现 education.Query 接口
func (h *Handler) List(ctx context.Context, listQuery educationDomain.ListQuery) ([]educationDomainViews.EducationView, int64, error) {
	var items []models.Education
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.EducationListQueryToCommonQuery(listQuery)

	queryBuilder := h.db.WithContext(ctx).Model(&models.Education{}).Where("deleted_at IS NULL")

	// Apply search
	if commonQuery.Search != "" {
		queryBuilder = queryBuilder.Where(
			"school ILIKE ? OR degree ILIKE ? OR major ILIKE ? OR field_of_study ILIKE ?",
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

	return slice.MapP(items, mapper.EducationModelToDomain), total, nil
}
