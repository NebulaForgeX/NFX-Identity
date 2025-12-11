package query

import (
	"context"
	"errors"
	educationAppQueries "nfxid/modules/auth/application/education/queries"
	educationDomainErrors "nfxid/modules/auth/domain/education/errors"
	educationDomainViews "nfxid/modules/auth/domain/education/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/slice"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type educationPGQuery struct {
	db *gorm.DB
}

func NewEducationPGQuery(db *gorm.DB) *educationPGQuery {
	return &educationPGQuery{db: db}
}

func (q *educationPGQuery) GetByID(ctx context.Context, id uuid.UUID) (educationDomainViews.EducationView, error) {
	var m models.Education
	if err := q.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return educationDomainViews.EducationView{}, educationDomainErrors.ErrEducationViewNotFound
		}
		return educationDomainViews.EducationView{}, err
	}
	return mapper.EducationModelToDomain(&m), nil
}

func (q *educationPGQuery) GetByProfileID(ctx context.Context, profileID uuid.UUID) ([]educationDomainViews.EducationView, error) {
	var items []models.Education
	if err := q.db.WithContext(ctx).
		Where("profile_id = ?", profileID).
		Where("deleted_at IS NULL").
		Order("start_date DESC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	return slice.MapP(items, mapper.EducationModelToDomain), nil
}

func (q *educationPGQuery) GetList(ctx context.Context, listQuery educationAppQueries.EducationListQuery) ([]educationDomainViews.EducationView, int64, error) {
	var items []models.Education
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.EducationListQueryToCommonQuery(listQuery)

	queryBuilder := q.db.WithContext(ctx).Model(&models.Education{}).Where("deleted_at IS NULL")

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

func (q *educationPGQuery) GetCount(ctx context.Context) (int64, error) {
	var count int64
	if err := q.db.WithContext(ctx).Model(&models.Education{}).Where("deleted_at IS NULL").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
