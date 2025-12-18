package list

import (
	"context"
	badgeDomain "nfxid/modules/auth/domain/badge"
	badgeDomainViews "nfxid/modules/auth/domain/badge/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/query"

	"gorm.io/gorm"
)

// Generic 获取 Badge 列表，实现 badgeDomain.List 接口
func (h *Handler) Generic(ctx context.Context, listQuery badgeDomain.ListQuery) ([]*badgeDomainViews.BadgeView, int64, error) {
	var items []models.Badge
	var total int64

	itemsResult, total, err := query.ExecuteQuery(
		ctx,
		h.db.WithContext(ctx).Model(&models.Badge{}),
		mapper.BadgeQueryToCommonQuery(listQuery),
		badgeQueryConfig,
		func(db *gorm.DB, data *[]models.Badge) error {
			return db.Find(data).Error
		},
	)
	if err != nil {
		return nil, 0, err
	}
	items = itemsResult

	// Convert to pointers
	result := make([]*badgeDomainViews.BadgeView, len(items))
	for i, item := range items {
		view := mapper.BadgeDBToDomain(&item)
		result[i] = &view
	}
	return result, total, nil
}

