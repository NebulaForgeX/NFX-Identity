package query

import (
	"context"
	"errors"
	occupationAppQueries "nfxid/modules/auth/application/profile_occupation/queries"
	occupationDomainErrors "nfxid/modules/auth/domain/profile_occupation/errors"
	occupationDomainViews "nfxid/modules/auth/domain/profile_occupation/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/slice"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type occupationPGQuery struct {
	db *gorm.DB
}

func NewOccupationPGQuery(db *gorm.DB) *occupationPGQuery {
	return &occupationPGQuery{db: db}
}

func (q *occupationPGQuery) GetByID(ctx context.Context, id uuid.UUID) (occupationDomainViews.OccupationView, error) {
	var m models.Occupation
	if err := q.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return occupationDomainViews.OccupationView{}, occupationDomainErrors.ErrOccupationViewNotFound
		}
		return occupationDomainViews.OccupationView{}, err
	}
	return mapper.OccupationModelToDomain(&m), nil
}

func (q *occupationPGQuery) GetByProfileID(ctx context.Context, profileID uuid.UUID) ([]occupationDomainViews.OccupationView, error) {
	var items []models.Occupation
	if err := q.db.WithContext(ctx).
		Where("profile_id = ?", profileID).
		Where("deleted_at IS NULL").
		Order("start_date DESC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	return slice.MapP(items, mapper.OccupationModelToDomain), nil
}

func (q *occupationPGQuery) GetList(ctx context.Context, listQuery occupationAppQueries.OccupationListQuery) ([]occupationDomainViews.OccupationView, int64, error) {
	var items []models.Occupation
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.OccupationListQueryToCommonQuery(listQuery)

	queryBuilder := q.db.WithContext(ctx).Model(&models.Occupation{}).Where("deleted_at IS NULL")

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

	return slice.MapP(items, mapper.OccupationModelToDomain), total, nil
}

func (q *occupationPGQuery) GetCount(ctx context.Context) (int64, error) {
	var count int64
	if err := q.db.WithContext(ctx).Model(&models.Occupation{}).Where("deleted_at IS NULL").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
