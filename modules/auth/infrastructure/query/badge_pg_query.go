package query

import (
	"context"
	"errors"
	badgeAppQueries "nebulaid/modules/auth/application/badge/queries"
	badgeDomainErrors "nebulaid/modules/auth/domain/badge/errors"
	badgeDomainViews "nebulaid/modules/auth/domain/badge/views"
	"nebulaid/modules/auth/infrastructure/query/mapper"
	"nebulaid/modules/auth/infrastructure/rdb/models"
	"nebulaid/pkgs/query"
	"nebulaid/pkgs/utils/slice"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var badgeQueryConfig = &query.QueryConfig{
	SearchConfig: &query.SearchConfig{
		Fields:   []string{"name", "description"},
		Operator: "ILIKE",
		Logic:    "OR",
	},
}

type badgePGQuery struct {
	db *gorm.DB
}

func NewBadgePGQuery(db *gorm.DB) *badgePGQuery {
	return &badgePGQuery{db: db}
}

func (q *badgePGQuery) GetByID(ctx context.Context, badgeID uuid.UUID) (badgeDomainViews.BadgeView, error) {
	var m models.Badge
	if err := q.db.WithContext(ctx).Where("id = ?", badgeID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return badgeDomainViews.BadgeView{}, badgeDomainErrors.ErrBadgeViewNotFound
		}
		return badgeDomainViews.BadgeView{}, err
	}
	return mapper.BadgeDBToDomain(&m), nil
}

func (q *badgePGQuery) GetByName(ctx context.Context, name string) (badgeDomainViews.BadgeView, error) {
	var m models.Badge
	if err := q.db.WithContext(ctx).Where("name = ?", name).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return badgeDomainViews.BadgeView{}, badgeDomainErrors.ErrBadgeViewNotFound
		}
		return badgeDomainViews.BadgeView{}, err
	}
	return mapper.BadgeDBToDomain(&m), nil
}

func (q *badgePGQuery) GetList(ctx context.Context, listQuery badgeAppQueries.BadgeListQuery) ([]badgeDomainViews.BadgeView, int64, error) {
	items, total, err := query.ExecuteQuery(
		ctx,
		q.db.WithContext(ctx).Model(&models.Badge{}),
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
