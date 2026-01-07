package list

import (
	"context"
	userDomain "nfxid/modules/auth/domain/user"
	userDomainViews "nfxid/modules/auth/domain/user/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/rdb/views"
	"nfxid/pkgs/query"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Generic 获取 User 列表，实现 userDomain.List 接口
func (h *Handler) Generic(ctx context.Context, listQuery userDomain.ListQuery) ([]*userDomainViews.UserView, int64, error) {
	var items []views.UserWithRoleView
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.UserListQueryToCommonQuery(listQuery)

	baseQuery := h.db.WithContext(ctx).Model(&views.UserWithRoleView{})

	// Use ExecuteQuery for search, pagination, and sorting
	itemsResult, total, err := query.ExecuteQuery(
		ctx,
		baseQuery,
		&commonQuery,
		userQueryConfig,
		func(db *gorm.DB, data *[]views.UserWithRoleView) error {
			// Apply default sorting if no sort is specified
			if len(commonQuery.Sorts) == 0 {
				db = db.Order("user_created_at DESC")
			}
			return db.Find(data).Error
		},
	)
	if err != nil {
		return nil, 0, err
	}
	items = itemsResult

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
