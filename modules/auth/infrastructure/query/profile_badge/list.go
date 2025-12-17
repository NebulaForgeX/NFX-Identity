package profile_badge

import (
	"context"
	profileBadgeDomain "nfxid/modules/auth/domain/profile_badge"
	profileBadgeDomainViews "nfxid/modules/auth/domain/profile_badge/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/slice"
)

// List 获取 ProfileBadge 列表，实现 profile_badge.Query 接口
func (h *Handler) List(ctx context.Context, listQuery profileBadgeDomain.ListQuery) ([]profileBadgeDomainViews.ProfileBadgeView, int64, error) {
	var items []models.ProfileBadge
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.ProfileBadgeListQueryToCommonQuery(listQuery)

	queryBuilder := h.db.WithContext(ctx).Model(&models.ProfileBadge{})

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
		queryBuilder = queryBuilder.Order("earned_at DESC")
	}

	// Execute query
	if err := queryBuilder.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return slice.MapP(items, mapper.ProfileBadgeModelToDomain), total, nil
}
