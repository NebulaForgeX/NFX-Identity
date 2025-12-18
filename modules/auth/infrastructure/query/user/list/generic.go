package list

import (
	"context"
	userDomain "nfxid/modules/auth/domain/user"
	userDomainViews "nfxid/modules/auth/domain/user/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/rdb/views"

	"github.com/google/uuid"
)

// Generic 获取 User 列表，实现 userDomain.List 接口
func (h *Handler) Generic(ctx context.Context, listQuery userDomain.ListQuery) ([]*userDomainViews.UserView, int64, error) {
	var items []views.UserWithRoleView
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.UserListQueryToCommonQuery(listQuery)

	queryBuilder := h.db.WithContext(ctx).Model(&views.UserWithRoleView{})

	// Apply search
	if commonQuery.Search != "" {
		queryBuilder = queryBuilder.Where(
			"username ILIKE ? OR email ILIKE ? OR phone ILIKE ?",
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
		queryBuilder = queryBuilder.Order("user_created_at DESC")
	}

	// Execute query
	if err := queryBuilder.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	// Convert to domain views
	result := make([]*userDomainViews.UserView, len(items))
	userIDs := make([]uuid.UUID, len(items))
	for i, item := range items {
		userIDs[i] = item.UserID
	}

	// Fetch all users in batch
	var users []models.User
	if err := h.db.WithContext(ctx).Where("id IN ?", userIDs).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	userMap := make(map[uuid.UUID]*models.User)
	for i := range users {
		userMap[users[i].ID] = &users[i]
	}

	for i, item := range items {
		u := userMap[item.UserID]
		view := mapper.UserViewToDomain(&item, u)
		result[i] = &view
	}

	return result, total, nil
}

