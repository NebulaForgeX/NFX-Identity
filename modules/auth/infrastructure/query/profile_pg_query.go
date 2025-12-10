package query

import (
	"context"
	"errors"
	profileAppQueries "nebulaid/modules/auth/application/profile/queries"
	profileDomainErrors "nebulaid/modules/auth/domain/profile/errors"
	profileDomainViews "nebulaid/modules/auth/domain/profile/views"
	"nebulaid/modules/auth/infrastructure/query/mapper"
	"nebulaid/modules/auth/infrastructure/rdb/models"
	"nebulaid/modules/auth/infrastructure/rdb/views"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type profilePGQuery struct {
	db *gorm.DB
}

func NewProfilePGQuery(db *gorm.DB) *profilePGQuery {
	return &profilePGQuery{db: db}
}

func (q *profilePGQuery) GetByID(ctx context.Context, id uuid.UUID) (profileDomainViews.ProfileView, error) {
	var v views.ProfileCompleteView
	if err := q.db.WithContext(ctx).Where("profile_id = ?", id).First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return profileDomainViews.ProfileView{}, profileDomainErrors.ErrProfileViewNotFound
		}
		return profileDomainViews.ProfileView{}, err
	}
	return mapper.ProfileViewToDomain(&v), nil
}

func (q *profilePGQuery) GetByUserID(ctx context.Context, userID uuid.UUID) (profileDomainViews.ProfileView, error) {
	var v views.ProfileCompleteView
	if err := q.db.WithContext(ctx).Where("user_id = ?", userID).First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return profileDomainViews.ProfileView{}, profileDomainErrors.ErrProfileViewNotFound
		}
		return profileDomainViews.ProfileView{}, err
	}
	return mapper.ProfileViewToDomain(&v), nil
}

func (q *profilePGQuery) GetList(ctx context.Context, listQuery profileAppQueries.ProfileListQuery) ([]profileDomainViews.ProfileView, int64, error) {
	var items []views.ProfileCompleteView
	var total int64

	listQuery.Normalize()
	commonQuery := mapper.ProfileListQueryToCommonQuery(listQuery)

	queryBuilder := q.db.WithContext(ctx).Model(&views.ProfileCompleteView{})

	// Apply search
	if commonQuery.Search != "" {
		queryBuilder = queryBuilder.Where(
			"display_name ILIKE ? OR nickname ILIKE ? OR bio ILIKE ?",
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
		queryBuilder = queryBuilder.Order("profile_created_at DESC")
	}

	// Execute query
	if err := queryBuilder.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	// Convert to domain views
	result := make([]profileDomainViews.ProfileView, len(items))
	for i, item := range items {
		result[i] = mapper.ProfileViewToDomain(&item)
	}

	return result, total, nil
}

func (q *profilePGQuery) GetCount(ctx context.Context) (int64, error) {
	var count int64
	if err := q.db.WithContext(ctx).Model(&models.Profile{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
