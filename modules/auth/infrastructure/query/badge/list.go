package badge

import (
	"context"
	badgeDomain "nfxid/modules/auth/domain/badge"
	badgeDomainViews "nfxid/modules/auth/domain/badge/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/query"
	"nfxid/pkgs/utils/slice"

	"gorm.io/gorm"
)

// List 获取 Badge 列表，实现 badge.Query 接口
func (h *Handler) List(ctx context.Context, listQuery badgeDomain.ListQuery) ([]badgeDomainViews.BadgeView, int64, error) {
	items, total, err := query.ExecuteQuery(
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
	return slice.MapP(items, mapper.BadgeDBToDomain), total, nil
}
