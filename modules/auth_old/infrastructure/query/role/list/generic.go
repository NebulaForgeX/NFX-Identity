package list

import (
	"context"
	roleDomain "nfxid/modules/auth/domain/role"
	roleDomainViews "nfxid/modules/auth/domain/role/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// Generic 获取 Role 列表，实现 roleDomain.List 接口
func (h *Handler) Generic(ctx context.Context, listQuery roleDomain.ListQuery) ([]*roleDomainViews.RoleView, int64, error) {
	var items []models.Role
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.RoleListQueryToCommonQuery(listQuery)

	queryBuilder := h.db.WithContext(ctx).Model(&models.Role{})

	// Apply search
	if commonQuery.Search != "" {
		queryBuilder = queryBuilder.Where(
			"name ILIKE ? OR description ILIKE ?",
			"%"+commonQuery.Search+"%",
			"%"+commonQuery.Search+"%",
		)
	}

	// Apply filters
	if listQuery.IsSystem != nil {
		queryBuilder = queryBuilder.Where("is_system = ?", *listQuery.IsSystem)
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
		queryBuilder = queryBuilder.Order("created_at DESC")
	}

	// Execute query
	if err := queryBuilder.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	// Convert to domain views
	result := make([]*roleDomainViews.RoleView, len(items))
	for i, item := range items {
		view := mapper.RoleModelToDomain(&item)
		result[i] = &view
	}

	return result, total, nil
}
