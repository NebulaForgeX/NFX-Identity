package list

import (
	"context"
	userRoleDomain "nfxid/modules/auth/domain/user_role"
	userRoleDomainViews "nfxid/modules/auth/domain/user_role/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/slice"
)

// Generic 获取 UserRole 列表，实现 userRoleDomain.List 接口
func (h *Handler) Generic(ctx context.Context, listQuery userRoleDomain.ListQuery) ([]*userRoleDomainViews.UserRoleView, int64, error) {
	var items []models.UserRole
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.UserRoleListQueryToCommonQuery(listQuery)

	queryBuilder := h.db.WithContext(ctx).Model(&models.UserRole{})

	// Apply filters
	if len(listQuery.UserIDs) > 0 {
		queryBuilder = queryBuilder.Where("user_id IN ?", listQuery.UserIDs)
	}
	if len(listQuery.RoleIDs) > 0 {
		queryBuilder = queryBuilder.Where("role_id IN ?", listQuery.RoleIDs)
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

	result := slice.MapP(items, mapper.UserRoleModelToDomain)
	// Convert to pointers
	pointerResult := make([]*userRoleDomainViews.UserRoleView, len(result))
	for i := range result {
		pointerResult[i] = &result[i]
	}
	return pointerResult, total, nil
}
